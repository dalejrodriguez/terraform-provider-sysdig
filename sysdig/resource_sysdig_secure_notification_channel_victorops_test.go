package sysdig_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/draios/terraform-provider-sysdig/sysdig"
)

func TestAccSecureNotificationChannelVictorOps(t *testing.T) {
	//var ncBefore, ncAfter secure.NotificationChannel

	rText := func() string { return acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum) }

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if v := os.Getenv("SYSDIG_SECURE_API_TOKEN"); v == "" {
				t.Fatal("SYSDIG_SECURE_API_TOKEN must be set for acceptance tests")
			}
		},
		Providers: map[string]terraform.ResourceProvider{
			"sysdig": sysdig.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: secureNotificationChannelVictorOpsWithName(rText()),
			},
		},
	})
}

func secureNotificationChannelVictorOpsWithName(name string) string {
	return fmt.Sprintf(`
resource "sysdig_secure_notification_channel_victorops" "sample-victorops" {
	name = "Example Channel %s - VictorOps"
	enabled = true
	api_key = "1234342-4234243-4234-2"
	routing_key = "My team"
	notify_when_ok = false
	notify_when_resolved = false
	send_test_notification = false
}`, name)
}
