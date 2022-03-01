package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var CurrentPid int

type Conf struct {
	Values []string `json:"values"`
}

func Rest() {
	router := gin.Default()

	router.GET("/job", GetJobs)
	router.POST("/job", RunJob)
	router.DELETE("/job", StopJob)

	err := router.Run()
	if err != nil {
		log.Error("Failed starting the service: ", err)
		return
	}
}

func GetJobs(c *gin.Context) {
	if CurrentPid > 0 {
		data, err := os.ReadFile("sites.txt")
		if err != nil {
			log.Error("Failed opening sites file: ", err)
			NewError(c, 500, fmt.Errorf("failed opening sites file: %w", err))
		}
		c.JSON(200, gin.H{"sites": string(data)})
	} else {
		c.JSON(404, gin.H{"sites": "none"})
	}
}

func RunJob(c *gin.Context) {
	conf := Conf{}

	err := c.BindJSON(&conf)
	if err != nil {
		log.Error("Failed parsing request JSON: ", err)
		NewError(c, 500, fmt.Errorf("failed parsing request JSON: %w", err))
		return
	}

	if pid, err := InstallDeployment(&conf); err != nil {
		log.Error("Failed installing a job: ", err)
		NewError(c, 500, fmt.Errorf("failed installing a job: %w", err))
	} else {
		CurrentPid = pid
		c.JSON(200, gin.H{"status": "started"})
	}
}

func InstallDeployment(conf *Conf) (int, error) {
	if len(conf.Values) == 0 {
		return -1, fmt.Errorf("empty values")
	}

	f, err := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return -1, err
	}

	_, err = f.Write([]byte(strings.Join(conf.Values, " ")))
	if err != nil {
		return -1, err
	}

	if err = f.Close(); err != nil {
		return -1, err
	}

	pid, err := start()
	if err != nil {
		return -1, err
	}

	return pid, nil
}

func StopJob(c *gin.Context) {
	err := kill(CurrentPid)
	if err != nil {
		log.Error("Failed killing a job: ", err)
		NewError(c, 500, fmt.Errorf("failed killing a job: %w", err))
	} else {
		CurrentPid = -1
		c.JSON(200, gin.H{"status": "killed"})
	}
}

func start() (int, error) {
	app := "sh"
	cmd := exec.Command(app, "start.sh")

	log.Info("executing command: ", cmd.String())
	err := cmd.Start()
	if err != nil {
		errExtended := fmt.Errorf("command execution failed: %w.", err)
		log.Error(errExtended)
		return 0, errExtended
	}
	pid := cmd.Process.Pid
	if err != nil {
		errExtended := fmt.Errorf("command execution failed: %w.", err)
		log.Error(errExtended)
		return 0, errExtended
	}
	log.Println(fmt.Sprintf("Started a process with pid: %d", pid))

	return pid, nil
}

func kill(pid int) error {
	app := "kill"
	cmd := exec.Command(app, "-9", strconv.Itoa(pid))
	log.Info("executing command: ", cmd.String())
	out, err := cmd.CombinedOutput()
	outFiltered := strings.Join(strings.Split(string(out), "\n"), "\n")
	if err != nil {
		errExtended := fmt.Errorf("command execution failed: %w.\n%s", err, outFiltered)
		log.Error(errExtended)
		return errExtended
	}
	return nil
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
