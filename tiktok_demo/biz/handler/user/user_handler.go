/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	user "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/model/basic/user"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/mw/jwt"
	service "github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/biz/service/user"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/errno"
	"github.com/cloudwego/hertz-examples/bizdemo/tiktok_demo/pkg/utils"
)

// UserRegister user registration api
//
// @router /douyin/user/register/  [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	_, err = service.NewUserService(ctx, c).UserRegister(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	jwt.JwtMiddleware.LoginHandler(ctx, c)
	token := c.GetString("token")
	v, _ := c.Get("user_id")
	user_id := v.(int64)
	c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     user_id,
	})
}

// UserLogin user login api
//
// @router /douyin/user/login/  [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get("user_id")
	user_id := v.(int64)
	token := c.GetString("token")
	c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     user_id,
	})
}

// User get user info
//
// @router /douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	u, err := service.NewUserService(ctx, c).UserInfo(&req)

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, user.DouyinUserResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		User:       u,
	})
}
