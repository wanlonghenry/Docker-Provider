// +build windows

package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/ingest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Microsoft/go-winio"
	"github.com/tinylib/msgp/msgp"
)

func CreateWindowsNamedPipesClient(namedPipe string) {
	containerLogPipePath := "\\\\.\\\\pipe\\\\" + namedPipe

	Log("Windows AMA::%s", containerLogPipePath)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	conn, err := winio.DialPipeAccess(ctx, containerLogPipePath, syscall.GENERIC_WRITE)

	if err != nil {
		Log("Error::Windows AMA::Unable to open Named Pipe %s", err.Error())
	} else {
		Log("Windows Named Pipe opened without any errors")
		ContainerLogNamedPipe = conn
	}

}
