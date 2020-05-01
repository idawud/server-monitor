package main
import (
	"fmt"
	"github.com/idawud/server-monitor/color-print"
	"github.com/idawud/server-monitor/data"
	"github.com/idawud/server-monitor/service"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

func cli() {

	writer := tabwriter.NewWriter(os.Stdout, 0, 16, 1, '\t', tabwriter.AlignRight)
	_, _ = fmt.Fprintln(writer, "Endpoint\t\t\tAvailability")
	_ = writer.Flush()

	for {
		printEndpointStatus(data.ENDPOINTS, writer)
		time.Sleep(time.Second * 25)
	}
}

func printEndpointStatus(endpoints []string, writer *tabwriter.Writer) {
	for _, endpoint := range endpoints {
		availability := service.CheckEndpointAvailable(endpoint)
		if availability {
			_, _ = fmt.Fprintln(writer, color_print.Green(endpoint+"\t\t\t"+strconv.FormatBool(availability)))
		} else {
			_, _ = fmt.Fprintln(writer, color_print.Red(endpoint+"\t\t\t"+strconv.FormatBool(availability)))
		}
		_ = writer.Flush()
	}
}

