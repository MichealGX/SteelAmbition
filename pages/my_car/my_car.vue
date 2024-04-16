<template>
  <view class="container">
    <view class="module core-module">
      <text class="module-title">核心模块</text>
      <text class="module-description">具备红外接收能力</text>
      <progress-bar :damage="coreModuleDamage"></progress-bar>
    </view>
    <view class="module weapon-module">
      <text class="module-title">武器模块</text>
      <text class="module-description">具备红外发射及接收功能</text>
      <progress-bar :damage="weaponModuleDamage"></progress-bar>
    </view>
    <view class="module aesthetic-module">
      <text class="module-title">外观模块</text>
      <text class="module-description">装饰作用</text>
      <!-- 外观模块可能不需要显示损坏程度 -->
    </view>
    <view class="module defense-module">
      <text class="module-title">防御模块</text>
      <text class="module-description">具备红外接收器</text>
      <progress-bar :damage="defenseModuleDamage"></progress-bar>
    </view>
    <view class="module mobility-module">
      <text class="module-title">行走模块</text>
      <text class="module-description">核心模块之一</text>
      <progress-bar :damage="mobilityModuleDamage"></progress-bar>
    </view>
  </view>
</template>

<script>
import ProgressBar from "@/components/ProgressBar.vue"; // 假设您有一个进度条组件

export default {
  components: {
    ProgressBar
  },
  data() {
    return {
      coreModuleDamage: 20, // 核心模块损坏程度，0为完好，100为完全损坏
      weaponModuleDamage: 35,
      defenseModuleDamage: 10,
      mobilityModuleDamage: 50,
    }
  },
  methods: {
    // 定义你的组件逻辑和方法
  }
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  background-color: #f5f5f5;
}

.module {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin: 10px;
  padding: 20px;
  border: 2px solid #000;
  border-radius: 10px;
  background-color: #ffffff;
  width: 80%;
  box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
}

.module-title {
  font-size: 20px;
  color: #333;
  font-weight: bold;
}

.module-description {
  font-size: 14px;
  color: #666;
  margin-bottom: 10px;
}
</style>
