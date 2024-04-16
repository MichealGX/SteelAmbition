<template>
  <view class="container">
    <view class="record" v-for="(record, index) in matchRecords" :key="index">
      <text class="record-date">{{ record.date }}</text>
      <text class="record-detail">用户1: {{ record.user }} vs 用户2: {{ record.opponent }} - 结果: {{ record.result }}</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      matchRecords: [
        { date: "2024-04-10", user: "玩家A", opponent: "玩家B", result: "胜利" },
        { date: "2024-04-09", user: "玩家A", opponent: "玩家C", result: "失败" },
        // 更多记录...
      ]
    };
  }
};
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}

.record {
  margin: 10px;
  padding: 15px;
  width: 90%;
  background-color: #f9f9f9;
  border: 1px solid #e1e1e1;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.record-date {
  font-size: 14px;
  margin-bottom: 5px;
  font-weight: bold;
}

.record-detail {
  font-size: 14px;
}
</style>
