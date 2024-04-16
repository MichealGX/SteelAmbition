<template>
  <view class="container">
    <form @submit.prevent="submitVehicle">
      <view class="input-group">
        <text>车辆名称:</text>
        <input type="text" v-model="vehicle.vehicle_name" placeholder="请输入车辆名称" required />
      </view>
      <view class="input-group">
        <text>核心模块重量值:</text>
        <input type="number" v-model.number="vehicle.core_w" placeholder="请输入重量值" required />
      </view>
      <view class="input-group">
        <text>核心模块能量值:</text>
        <input type="number" v-model.number="vehicle.core_e" placeholder="请输入能量值" required />
      </view>
      <view class="input-group">
        <text>外观模块重量值:</text>
        <input type="number" v-model.number="vehicle.appearance_w" placeholder="请输入重量值" />
      </view>
      <view class="input-group">
        <text>武器模块重量值:</text>
        <input type="number" v-model.number="vehicle.weapon_w" placeholder="请输入重量值" required />
      </view>
      <view class="input-group">
        <text>武器模块能量值:</text>
        <input type="number" v-model.number="vehicle.weapon_e" placeholder="请输入能量值" required />
      </view>
      <view class="input-group">
        <text>防御模块重量值:</text>
        <input type="number" v-model.number="vehicle.defence_w" placeholder="请输入重量值" />
      </view>
      <view class="input-group">
        <text>防御模块能量值:</text>
        <input type="number" v-model.number="vehicle.defence_e" placeholder="请输入能量值" />
      </view>
      <view class="input-group">
        <text>行走模块重量值:</text>
        <input type="number" v-model.number="vehicle.walking_w" placeholder="请输入重量值" required />
      </view>
      <view class="input-group">
        <text>行走模块能量值:</text>
        <input type="number" v-model.number="vehicle.walking_e" placeholder="请输入能量值" required />
      </view>
      <view class="input-group">
        <text>行走模块速度值:</text>
        <input type="number" v-model.number="vehicle.walking_s" placeholder="请输入速度值" required />
      </view>
      <button type="submit">提交车辆信息</button>
    </form>
  </view>
</template>

<script>
export default {
  data() {
    return {
      vehicle: {
        vehicle_name: '',
        core_w: 0,
        core_e: 0,
        appearance_w: null, // 可选字段为空时不提交
        weapon_w: 0,
        weapon_e: 0,
        defence_w: null, // 可选字段为空时不提交
        defence_e: null, // 可选字段为空时不提交
        walking_w: 0,
        walking_e: 0,
        walking_s: 0,
      }
    };
  },
  methods: {
    async submitVehicle() {
      // 过滤掉为空的可选字段
      const payload = Object.fromEntries(Object.entries(this.vehicle).filter(([_, v]) => v !== null));
      try {
        const response = await uni.request({
          url: 'http://127.0.0.1:4523/m1/4225974-0-default/platform/vehicles/add',
          method: 'POST',
          data: payload,
          header: {
            'content-type': 'application/json'
          },
        });
        if (response[1].statusCode === 200) {
          uni.showToast({
            title: '车辆添加成功',
            icon: 'success',
            duration: 2000
          });
          // Optionally reset form or navigate away
        } else {
          uni.showToast({
            title: '添加失败',
            icon: 'none',
            duration: 2000
          });
        }
      } catch (error) {
        console.error("Error submitting vehicle:", error);
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
