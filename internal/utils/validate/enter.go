package validate

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
)

var trans ut.Translator

/*
翻译器初始化，init 在包被引入时自动执行，
*/
func init() {
	// 创建翻译器
	zhLocale := zh.New()
	uni := ut.New(zhLocale, zhLocale)
	var found bool
	trans, found = uni.GetTranslator("zh")
	if !found {
		logrus.Warnf("获取中文翻译器失败，将默认使用英文校验提示")
		trans = nil
		return
	}

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		logrus.Warnf("获取Validator 失败，参数校验无法正常进行")
		trans = nil
		return
	}

	err := zh_translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		logrus.Warnf("注册中文翻译失败: %v，默认使用英文校验提示", err)
		trans = nil
		return
	}

	// 注册自定义字段名解析(优先使用label标签)
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		return label
	})
}

/*
验证错误转化为中文
*/
func ValidateError(err error) string {
	// 如果翻译器没初始化成功，降级处理
	if trans == nil {
		return err.Error()
	}
	var errs validator.ValidationErrors
	// 使用 errors.As 穿透错误包装，更加健壮
	if !errors.As(err, &errs) {
		return err.Error()
	}
	// 预分配切片容量
	list := make([]string, 0, len(errs))
	for _, e := range errs {
		msg := e.Translate(trans)
		if msg == "" {
			msg = e.Error() //如果翻译不到，就兜底用英文或原始错误
		}
		list = append(list, msg)
	}
	return strings.Join(list, ";")
}
