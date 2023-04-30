// Code generated by github.com/willdhorn/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/willdhorn/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_ValidType_validArgs_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["break"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("break"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["break"] = arg0
	var arg1 string
	if tmp, ok := rawArgs["default"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("default"))
		arg1, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["default"] = arg1
	var arg2 string
	if tmp, ok := rawArgs["func"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("func"))
		arg2, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["func"] = arg2
	var arg3 string
	if tmp, ok := rawArgs["interface"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("interface"))
		arg3, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["interface"] = arg3
	var arg4 string
	if tmp, ok := rawArgs["select"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("select"))
		arg4, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["select"] = arg4
	var arg5 string
	if tmp, ok := rawArgs["case"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("case"))
		arg5, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["case"] = arg5
	var arg6 string
	if tmp, ok := rawArgs["defer"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("defer"))
		arg6, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["defer"] = arg6
	var arg7 string
	if tmp, ok := rawArgs["go"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("go"))
		arg7, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["go"] = arg7
	var arg8 string
	if tmp, ok := rawArgs["map"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("map"))
		arg8, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["map"] = arg8
	var arg9 string
	if tmp, ok := rawArgs["struct"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("struct"))
		arg9, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["struct"] = arg9
	var arg10 string
	if tmp, ok := rawArgs["chan"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("chan"))
		arg10, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["chan"] = arg10
	var arg11 string
	if tmp, ok := rawArgs["else"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("else"))
		arg11, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["else"] = arg11
	var arg12 string
	if tmp, ok := rawArgs["goto"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("goto"))
		arg12, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["goto"] = arg12
	var arg13 string
	if tmp, ok := rawArgs["package"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
		arg13, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["package"] = arg13
	var arg14 string
	if tmp, ok := rawArgs["switch"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("switch"))
		arg14, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["switch"] = arg14
	var arg15 string
	if tmp, ok := rawArgs["const"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("const"))
		arg15, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["const"] = arg15
	var arg16 string
	if tmp, ok := rawArgs["fallthrough"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("fallthrough"))
		arg16, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["fallthrough"] = arg16
	var arg17 string
	if tmp, ok := rawArgs["if"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("if"))
		arg17, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["if"] = arg17
	var arg18 string
	if tmp, ok := rawArgs["range"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("range"))
		arg18, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["range"] = arg18
	var arg19 string
	if tmp, ok := rawArgs["type"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
		arg19, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["type"] = arg19
	var arg20 string
	if tmp, ok := rawArgs["continue"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("continue"))
		arg20, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["continue"] = arg20
	var arg21 string
	if tmp, ok := rawArgs["for"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("for"))
		arg21, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["for"] = arg21
	var arg22 string
	if tmp, ok := rawArgs["import"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("import"))
		arg22, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["import"] = arg22
	var arg23 string
	if tmp, ok := rawArgs["return"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("return"))
		arg23, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["return"] = arg23
	var arg24 string
	if tmp, ok := rawArgs["var"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("var"))
		arg24, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["var"] = arg24
	var arg25 string
	if tmp, ok := rawArgs["_"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("_"))
		arg25, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["_"] = arg25
	return args, nil
}

func (ec *executionContext) field_ValidType_validInputKeywords_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 *ValidInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalOValidInput2ᚖgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐValidInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Content_Post_foo(ctx context.Context, field graphql.CollectedField, obj *ContentPost) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Content_Post_foo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Content_Post_foo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Content_Post",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Content_User_foo(ctx context.Context, field graphql.CollectedField, obj *ContentUser) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Content_User_foo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Foo, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Content_User_foo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Content_User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ValidType_differentCase(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ValidType_differentCase(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DifferentCase, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ValidType_differentCase(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ValidType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ValidType_different_case(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ValidType_different_case(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DifferentCaseOld, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ValidType_different_case(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ValidType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _ValidType_validInputKeywords(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ValidType_validInputKeywords(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ValidInputKeywords, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ValidType_validInputKeywords(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ValidType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_ValidType_validInputKeywords_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _ValidType_validArgs(ctx context.Context, field graphql.CollectedField, obj *ValidType) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_ValidType_validArgs(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ValidArgs, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_ValidType_validArgs(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "ValidType",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_ValidType_validArgs_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputValidInput(ctx context.Context, obj interface{}) (ValidInput, error) {
	var it ValidInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "for", "import", "return", "var", "_"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "break":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("break"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Break = data
		case "default":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("default"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Default = data
		case "func":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("func"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Func = data
		case "interface":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("interface"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Interface = data
		case "select":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("select"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Select = data
		case "case":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("case"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Case = data
		case "defer":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("defer"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Defer = data
		case "go":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("go"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Go = data
		case "map":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("map"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Map = data
		case "struct":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("struct"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Struct = data
		case "chan":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("chan"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Chan = data
		case "else":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("else"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Else = data
		case "goto":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("goto"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Goto = data
		case "package":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Package = data
		case "switch":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("switch"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Switch = data
		case "const":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("const"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Const = data
		case "fallthrough":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("fallthrough"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Fallthrough = data
		case "if":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("if"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.If = data
		case "range":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("range"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Range = data
		case "type":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Type = data
		case "continue":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("continue"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Continue = data
		case "for":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("for"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.For = data
		case "import":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("import"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Import = data
		case "return":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("return"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Return = data
		case "var":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("var"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Var = data
		case "_":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("_"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Underscore = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Content_Child(ctx context.Context, sel ast.SelectionSet, obj ContentChild) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case ContentUser:
		return ec._Content_User(ctx, sel, &obj)
	case *ContentUser:
		if obj == nil {
			return graphql.Null
		}
		return ec._Content_User(ctx, sel, obj)
	case ContentPost:
		return ec._Content_Post(ctx, sel, &obj)
	case *ContentPost:
		if obj == nil {
			return graphql.Null
		}
		return ec._Content_Post(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var content_PostImplementors = []string{"Content_Post", "Content_Child"}

func (ec *executionContext) _Content_Post(ctx context.Context, sel ast.SelectionSet, obj *ContentPost) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, content_PostImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Content_Post")
		case "foo":

			out.Values[i] = ec._Content_Post_foo(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var content_UserImplementors = []string{"Content_User", "Content_Child"}

func (ec *executionContext) _Content_User(ctx context.Context, sel ast.SelectionSet, obj *ContentUser) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, content_UserImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Content_User")
		case "foo":

			out.Values[i] = ec._Content_User_foo(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var validTypeImplementors = []string{"ValidType"}

func (ec *executionContext) _ValidType(ctx context.Context, sel ast.SelectionSet, obj *ValidType) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, validTypeImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ValidType")
		case "differentCase":

			out.Values[i] = ec._ValidType_differentCase(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "different_case":

			out.Values[i] = ec._ValidType_different_case(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "validInputKeywords":

			out.Values[i] = ec._ValidType_validInputKeywords(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "validArgs":

			out.Values[i] = ec._ValidType_validArgs(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalOValidInput2ᚖgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐValidInput(ctx context.Context, v interface{}) (*ValidInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputValidInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOValidType2ᚖgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐValidType(ctx context.Context, sel ast.SelectionSet, v *ValidType) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._ValidType(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
