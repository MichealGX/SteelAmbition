<template>
  <view class="progress-bar">
    <view class="progress" :style="{ width: progressWidth }"></view>
  </view>
</template>

<script>
export default {
  props: {
    damage: {
      type: Number,
      default: 0
    }
  },
  computed: {
    progressWidth() {
      const width = 100 - this.damage;
      return `${width}%`;
    }
  }
}
</script>

<style scoped>
.progress-bar {
  width: 100%;
  height: 20px;
  background-color: #e5e5e5;
  border-radius: 10px;
  overflow: hidden;
}

.progress {
  height: 100%;
  background-color: #f00;
  transition: width 0.3s ease-in-out;
}
</style>
