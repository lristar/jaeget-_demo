package pkg

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var errTag = "reg_error_info"

type User struct {
	Name string   `validate:"min=6,max=10" reg_error_info:"名字格式有问题"`
	Age  int      `validate:"min=1,max=100" reg_error_info:"年龄格式有问题"`
	AA   []string `validate:"gt=0" reg_error_info:"AA不能为空"`
	Aa   *Person   `validate:"required,dive" reg_error_info:"GG不能为空"`
}

type Person struct {
	Msg string `validate:"required" reg_error_info:"msg格式有问题"`
	//nums []string `validate:"gt=0" reg_error_info:"nums不能为空"`
}

func ValidateStruct() {
	validate := validator.New()
	u2 := &User{
		Name: "adfadfd",
		Age:  15,
		AA:   []string{""},
		Aa:   &Person{},
	}
	err := validate.Struct(u2)
	fmt.Println("err is ", err)
	fmt.Println(processErr(u2, err))
}

//func RegisterVal(v *validator.Validate) {
//	for tag, f := range checkFunc {
//		_ = v.RegisterValidation(tag, f)
//	}
//}

// tag 返回错误自定义

func processErr(u interface{}, err error) string {
	if err == nil { //如果为nil 阐明校验通过
		return ""
	}
	// 用一个map存结构体
	invalid, ok := err.(*validator.InvalidValidationError) //如果是输出参数有效，则间接返回输出参数谬误
	if ok {
		return "输出参数谬误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Namespace() //获取是哪个字段不合乎格局
		// 获取的结构体需要往下判断
		errorInfo, ok := FindTag(fieldName,u) //通过反射获取filed
		if ok {
			return fieldName + ":" + errorInfo           //返回谬误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}

// 将结构体内的结构体都保存一下
//func ForeachStruct(u interface{}) {
//	param := make(map[string]interface{}, 0)
//	temp := reflect.ValueOf(u)
//	types := temp.Type().Elem()
//	values := temp.Elem()
//	for i := 0; i < types.NumField(); i++ {
//		if types.Field(i).Type.Kind() == reflect.Struct {
//			param[types.Field(i).Name] = values.Field(
//		}
//	}
//	fmt.Println("param is :", param)
//}

func entryValue(value reflect.Value) error{
	if value.Kind()!=reflect.Ptr {
		return errors.New("不是结构体")
	}
	elemType, elemValue := value.Type().Elem(), value.Elem()
	for i := 0; i < elemType.NumField(); i++ {
		if !elemValue.Field(i).CanSet() {
			continue
		}
		if elemValue.Field(i).Kind() == reflect.Ptr{
			entryValue(elemValue.Field(i))
		}
		fieldType := elemType.Field(i)
			if elemValue.Field(i).IsZero() {
				tag := fieldType.Tag.Get(errTag)
				fmt.Println(tag)
			} else {
				fmt.Println("非空", fieldType)
			}
		}
	return nil
	}

// 根据方法去查找tag User.Aa.Msg
func FindTag(name string,u interface{})(string,bool)  {
	var (
		rt reflect.StructField
	    ok bool
	)
	sp :=strings.Split(name,".")
	rf :=reflect.ValueOf(u)
	if rf.Kind() != reflect.Ptr{
		return "",false
	}
	for i:=1; i< len(sp); i++ {
		rt, ok =rf.Type().Elem().FieldByName(sp[i])
		if ok {
			switch rt.Type.Kind() {
			case reflect.Ptr:
				if !rf.Elem().FieldByName(sp[i]).IsZero() {
					rf = rf.Elem().FieldByName(sp[i])
				}else{
					return "",false
				}
			default:
				break
			}
		}
	}
	return rt.Tag.Get(errTag),true
}

//
