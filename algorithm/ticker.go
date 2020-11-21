package algortm

import (
	"context"
	"fmt"
	"time"
)

func timeT() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.println("11")
		case ctx.Done():
			fmt.println("done")
		}
	}
}
