<template>
	<view>
    <view style="width: 100%;height: 80px;display: flex;align-items: center;justify-content: space-around;">
       <button @click="showLogin" :class="{ active: activeTab === 'login' }">登录</button>
		   <button @click="showRegister" :class="{ active: activeTab === 'register' }">注册</button>
    </view>
    <view style="padding: 20px;">
        <view v-if="activeTab === 'login'">
		    	<input type="text" style="width: 100%;height: 45px;border-bottom: 1px solid #ccc;" v-model="loginData.username" placeholder="用户名">
		    	<input type="password"  style="width: 100%;height: 45px;border-bottom: 1px solid #ccc;" v-model="loginData.password" placeholder="密码">
          <button style="margin-top: 15px;" @click="logins">登录</button>
		    </view>
		    <view v-if="activeTab === 'register'">
		    	<input type="text"  style="width: 100%;height: 45px;border-bottom: 1px solid #ccc;" v-model="registerData.username" placeholder="用户名">
		    	<input type="password"  style="width: 100%;height: 45px;border-bottom: 1px solid #ccc;" v-model="registerData.password" placeholder="密码">
		    	<input type="emit"  style="width: 100%;height: 45px;border-bottom: 1px solid #ccc;" v-model="registerData.emit" placeholder="邮箱">
          <button style="margin-top: 15px;" @click="registers">注册</button>
		    </view>
    </view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				activeTab: 'login',
				loginData: {
					username: '',
					password: ''
				},
				registerData: {
					username: '',
					password: '',
					emit: ''
				}
			}
		},
		methods: {
			showLogin() {
				this.activeTab = 'login';
			},
			showRegister() {
				this.activeTab = 'register';
			},
      // 点击登录按钮
      logins(){
        let thit = this
        if(!thit.loginData.username){
          return
        }
        if(!thit.loginData.password){
          return
        }
         uni.request(`http://127.0.0.1:4523/m1/4225974-0-default/platform/users/login`, {
									method: 'POST',
									data: {
										username: thit.loginData.username || '',
										psw: thit.loginData.password || '',
									},
								}).then((res) => {
									// 处理响应数据
									console.log(res,'res')
                  if(res.code===200){
                    uni.switchTab({
                        url: '/pages/tabbar/index/index'
                     });
                  }
								}).catch((err) => {
									// 处理错误
									console.error(err)
								});
      },
      // 点击注册按钮
      registers(){
        let thit_ = this
        if(!thit_.registerData.username){
          return
        }
        if(!thit_.registerData.password){
          return
        }
         uni.request(`http://127.0.0.1:4523/m1/4225974-0-default/platform/users/register`, {
									method: 'POST',
									data: {
										username: thit_.registerData.username || '',
										password: thit_.registerData.password || '',
										email: thit_.registerData.emit ? thit_.registerData.emit : '',
									},
								}).then((res) => {
									// 处理响应数据
									console.log(res,'res')
                  if(res.code===200){
                    uni.navigateBack({ delta: 1 })
                  }
								}).catch((err) => {
									// 处理错误
									console.error(err)
								});
      }
		}
	}
</script>

<style>
	.active {
		color: #2e8726;
	}
	.view {
		text-align: center;
	}
  input{
    margin-top: 15px;
  }
</style>