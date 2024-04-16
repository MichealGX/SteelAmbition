<template>
  <view class="my-page">
    <view class="avatar-container">
      <!-- 这里是上传头像的圆形 -->
      <image class="avatar" src="/static/avatar_placeholder.png" mode="aspectFill"></image>
    </view>
    <view class="menu-list">
      <view class="menu-item" @click="goToLogin">
        <text>登录</text>
      </view>
      <view class="menu-item" @click="goToMyCar">
        <text>我的小车</text>
      </view>
      <view class="menu-item" @click="goToConfigCar">
        <text>配置小车</text>
      </view>
      <view class="menu-item" @click="goToGameRecords">
        <text>对局记录</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  methods: {
    goToLogin() {
      // 跳转到登录页面的逻辑
      uni.navigateTo({
        url: '/pages/login/login'
      });
    },
    goToMyCar() {
      // 跳转到我的小车页面的逻辑
      uni.navigateTo({
        url: '/pages/my_car/my_car'
      });
    },
    goToConfigCar() {
      // 跳转到配置小车页面的逻辑
      uni.navigateTo({
        url: '/pages/config_car/config_car'
      });
    },
    goToGameRecords() {
      // 跳转到对局记录页面的逻辑
      uni.navigateTo({
        url: '/pages/game_records/game_records'
      });
    }
  }
};
</script>

<style>
.my-page {
  padding: 20px;
}

.avatar-container {
  text-align: center;
}

.avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
}

.menu-list {
  margin-top: 20px;
}

.menu-item {
  padding: 10px;
  border-bottom: 1px solid #ccc;
  cursor: pointer;
}

.menu-item:last-child {
  border-bottom: none;
}
</style>
