<template>
  <view class="container">
    <view class="vehicles-display">
      <view v-for="(vehicle, index) in vehicles" :key="index" class="vehicle-card">
        <text class="vehicle-name">{{ vehicle.vehicle_name }}</text>
        <!-- 此处可展示更多车辆信息 -->
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      vehicles: []
    };
  },
  onShow() {
    this.getVehicles();
  },
  methods: {
    async getVehicles() {
      try {
        const response = await uni.request({
          url: 'http://127.0.0.1:4523/m1/4225974-0-default/platform/vehicles', // 这里应改为获取车辆列表的正确接口
          method: 'GET',
          header: {
            'content-type': 'application/json'
          },
        });
        if (response[1].statusCode === 200) {
          this.vehicles = response[1].data; // 假设后端直接返回车辆列表数组
        } else {
          uni.showToast({
            title: '获取车辆列表失败',
            icon: 'none',
            duration: 2000
          });
        }
      } catch (error) {
        console.error("Error getting vehicles:", error);
        uni.showToast({
          title: '网络错误',
          icon: 'none',
          duration: 2000
        });
      }
    },
  }
};
</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}
.vehicles-display {
  width: 100%;
}
.vehicle-card {
  padding: 10px;
  margin: 5px 0;
  border: 1px solid #ddd;
  border-radius: 5px;
}
.vehicle-name {
  font-weight: bold;
}
</style>
