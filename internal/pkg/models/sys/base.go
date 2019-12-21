package sys

import(
	"fmt"

	"github.com/liuning3399/app_go/internal/pkg/models/basemodel"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", basemodel.GetTablePrefix(),"sys_", name)
}