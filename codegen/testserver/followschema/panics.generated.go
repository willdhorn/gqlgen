// Code generated by github.com/willdhorn/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/willdhorn/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type PanicsResolver interface {
	FieldScalarMarshal(ctx context.Context, obj *Panics) ([]MarshalPanic, error)

	ArgUnmarshal(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Panics_argUnmarshal_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []MarshalPanic
	if tmp, ok := rawArgs["u"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("u"))
		arg0, err = ec.unmarshalNMarshalPanic2ᚕgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanicᚄ(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["u"] = arg0
	return args, nil
}

func (ec *executionContext) field_Panics_fieldFuncMarshal_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 []MarshalPanic
	if tmp, ok := rawArgs["u"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("u"))
		arg0, err = ec.unmarshalNMarshalPanic2ᚕgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanicᚄ(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["u"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Panics_fieldScalarMarshal(ctx context.Context, field graphql.CollectedField, obj *Panics) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Panics_fieldScalarMarshal(ctx, field)
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
		return ec.resolvers.Panics().FieldScalarMarshal(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]MarshalPanic)
	fc.Result = res
	return ec.marshalNMarshalPanic2ᚕgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanicᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Panics_fieldScalarMarshal(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Panics",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type MarshalPanic does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Panics_fieldFuncMarshal(ctx context.Context, field graphql.CollectedField, obj *Panics) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Panics_fieldFuncMarshal(ctx, field)
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
		return obj.FieldFuncMarshal(ctx, fc.Args["u"].([]MarshalPanic)), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]MarshalPanic)
	fc.Result = res
	return ec.marshalNMarshalPanic2ᚕgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanicᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Panics_fieldFuncMarshal(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Panics",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type MarshalPanic does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Panics_fieldFuncMarshal_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

func (ec *executionContext) _Panics_argUnmarshal(ctx context.Context, field graphql.CollectedField, obj *Panics) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Panics_argUnmarshal(ctx, field)
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
		return ec.resolvers.Panics().ArgUnmarshal(rctx, obj, fc.Args["u"].([]MarshalPanic))
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

func (ec *executionContext) fieldContext_Panics_argUnmarshal(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Panics",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
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
	if fc.Args, err = ec.field_Panics_argUnmarshal_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var panicsImplementors = []string{"Panics"}

func (ec *executionContext) _Panics(ctx context.Context, sel ast.SelectionSet, obj *Panics) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, panicsImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Panics")
		case "fieldScalarMarshal":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Panics_fieldScalarMarshal(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "fieldFuncMarshal":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Panics_fieldFuncMarshal(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "argUnmarshal":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Panics_argUnmarshal(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
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

func (ec *executionContext) unmarshalNMarshalPanic2githubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanic(ctx context.Context, v interface{}) (MarshalPanic, error) {
	var res MarshalPanic
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNMarshalPanic2githubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanic(ctx context.Context, sel ast.SelectionSet, v MarshalPanic) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNMarshalPanic2ᚕgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanicᚄ(ctx context.Context, v interface{}) ([]MarshalPanic, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]MarshalPanic, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNMarshalPanic2githubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanic(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNMarshalPanic2ᚕgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanicᚄ(ctx context.Context, sel ast.SelectionSet, v []MarshalPanic) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNMarshalPanic2githubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐMarshalPanic(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalOPanics2ᚖgithubᚗcomᚋwilldhornᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPanics(ctx context.Context, sel ast.SelectionSet, v *Panics) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Panics(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
