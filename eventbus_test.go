/**
 * @Author: kwens
 * @Date: 2023-12-05 15:53:08
 * @Description:
 */
package eventbus

import (
	"context"
	"fmt"
	"testing"

	"github.com/kwens/eventbus/val"
)

var UserEvent val.EventCallBack = func(ctx context.Context, data interface{}) error {
	fmt.Println("UserEvent Doing...")
	fmt.Printf("data is:%+v", data)
	return nil
}

func TestEventBus(t *testing.T) {
	// init
	eb := New()

	// subscirbe
	if err := eb.Subscribe("user.login", UserEvent); err != nil {
		t.Fatal(err)
	}

	// publish
	if err := eb.Publish(context.Background(), "user.login", "hello world"); err!= nil {
		t.Fatal(err)
	}
}
