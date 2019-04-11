package pythonanywhere

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"time"
)

// Client represents the client struct
type Client struct {
	client *http.Client
	base   string
	token  string
}

// NewClientWith represents the newclientwith command
func NewClientWith(username string, api_token string) (*Client, error) {
	url := "https://www.pythonanywhere.com/api/v0/user/" + username
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url+"/webapps/"+username+".pythonanywhere.com/reload/", nil)
	req.Header.Set("Authorization", "Token "+api_token)
	if err != nil {
		return nil, errors.New("error on response. Change the API Token")
	}
	response, err := netClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return &Client{netClient, url, api_token}, nil
}

// CreateWebapp represents the create webapp command
func (c *Client) CreateWebapp(domainName string, pythonVersion string) (string, error) {
	data := url.Values{}
	data.Set("domain_name", domainName)
	data.Add("python_version", pythonVersion)
	req, err := http.NewRequest("POST", c.base+"/webapps/", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		return "Error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// ListWebapps represents the list webapps command
func (c *Client) ListWebapps() (string, error) {
	req, err := http.NewRequest("GET", c.base+"/webapps/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// DeleteWebapp represents the delete webapp command
func (c *Client) DeleteWebapp(domainName string) (string, error) {
	req, err := http.NewRequest("DELETE", c.base+"/webapps/"+domainName+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// GetWebappInfo represents the get webapp info command
func (c *Client) GetWebappInfo(domainName string) (string, error) {
	req, err := http.NewRequest("GET", c.base+"/webapps/"+domainName+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// CreateConsole represents the create console command
func (c *Client) CreateConsole(executable string, arguments string, workingDirectory string) (string, error) {
	data := url.Values{}
	data.Set("executable", executable)
	data.Add("arguments", arguments)
	data.Add("working_directory", workingDirectory)

	req, err := http.NewRequest("POST", c.base+"/consoles/", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// ListConsoles represents the list consoles command
func (c *Client) ListConsoles() (string, error) {
	req, err := http.NewRequest("GET", c.base+"/consoles/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// ListSharedConsoles represents the list shared consoles command
func (c *Client) ListSharedConsoles() (string, error) {
	req, err := http.NewRequest("GET", c.base+"/consoles/shared_with_you/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// KillConsole represents the kill console command
func (c *Client) KillConsole(id string) (string, error) {
	req, err := http.NewRequest("DELETE", c.base+"/consoles/"+id+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// GetConsoleInfo represents the get console info command
func (c *Client) GetConsoleInfo(id string) (string, error) {
	req, err := http.NewRequest("GET", c.base+"/consoles/"+id+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// ShareFile represents the share file command
func (c *Client) ShareFile(path string) (string, error) {
	data := url.Values{}
	data.Set("path", path)
	req, err := http.NewRequest("POST", c.base+"/files/sharing/", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		return "Error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// CheckSharingStatus represents the check sharing status command
func (c *Client) CheckSharingStatus(path string) (string, error) {
	req, err := http.NewRequest("GET", c.base+"/files/tree/?path="+path, nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// StopSharing represents the stop sharing command
func (c *Client) StopSharing(path string) (string, error) {
	req, err := http.NewRequest("DELETE", c.base+"/files/sharing/?path="+path, nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// ListFolder represents the list folder command
func (c *Client) ListFolder(path string) (string, error) {
	req, err := http.NewRequest("GET", c.base+"/files/tree/?path="+path, nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// DeleteFile represents the delete file command
func (c *Client) DeleteFile(path string) (string, error) {
	req, err := http.NewRequest("DELETE", c.base+"/files/path"+path+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// CreateScheduledTask represents the create scheduled task command
func (c *Client) CreateScheduledTask(command string, enabled string, interval string, hour string, minute string) (string, error) {
	data := url.Values{}
	data.Set("enabled", enabled)
	data.Add("command", command)
	data.Add("interval", interval)
	data.Add("hour", hour)
	data.Add("minute", minute)

	req, err := http.NewRequest("POST", c.base+"/schedule/", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// ListScheduledTasks represents the list scheduled tasks command
func (c *Client) ListScheduledTasks() (string, error) {
	req, err := http.NewRequest("GET", c.base+"/schedule/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// GetScheduleTaskInfo returns information about a scheduled task
func (c *Client) GetScheduleTaskInfo(id string) (string, error) {
	req, err := http.NewRequest("GET", c.base+"/schedule/"+id+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}

// DeleteScheduledTask represents the delete scheduled task command
func (c *Client) DeleteScheduledTask(id string) (string, error) {
	req, err := http.NewRequest("DELETE", c.base+"/schedule/"+id+"/", nil)
	req.Header.Set("Authorization", "Token "+c.token)
	if err != nil {
		return "error on response. Change the API Token", err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return "Error on response", err
	}
	defer response.Body.Close()

	return "", nil
}
