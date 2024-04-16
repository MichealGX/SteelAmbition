<template>
  <view>
    <button @click="showVehicleSelection('create')">创建对战</button>
    <button @click="showVehicleSelection('join')">加入对战</button>

    <view class="uni-title uni-common-mt uni-common-pl">选择你的车辆</view>
    <view class="uni-list">
      <radio-group @change="radioChange">
        <label class="uni-list-cell uni-list-cell-pd" v-for="(item, index) in items" :key="item.value">
          <view>
            <radio :value="item.value" :checked="index === current" />
          </view>
          <view>{{item.name}}</view>
        </label>
      </radio-group>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      items: [
        { value: '1', name: 'A', checked: 'true' },
        { value: '2', name: 'B' },
        { value: '3', name: 'C' },
        { value: '4', name: 'D' },
        { value: '5', name: 'E' },
        { value: '6', name: 'F' },
      ],
      current: 0,
      selectedVehicle: null
    };
  },
  methods: {
    radioChange(evt) {
      for (let i = 0; i < this.items.length; i++) {
        if (this.items[i].value === evt.detail.value) {
          this.current = i;
          this.selectedVehicle = this.items[i];
          break;
        }
      }
    },
    showVehicleSelection(type) {
      if (type === 'create') {
        this.createGame();
      } else if (type === 'join') {
        this.joinGame();
      }
    },
    createGame() {
      // 处理创建对战逻辑
      // 将选择的车辆信息传递到游戏页面
      uni.navigateTo({
        url: '/pages/tabbar/gamecenter/game/game?selectedVehicle=' + JSON.stringify(this.selectedVehicle),
      });
    },
    joinGame() {
      // 处理加入对战逻辑
      // 将选择的车辆信息传递到游戏页面
      uni.navigateTo({
        url: '/pages/tabbar/gamecenter/game/game?selectedVehicle=' + JSON.stringify(this.selectedVehicle),
      });
    }
  },
};
</script>

<style>
/* 样式 */
</style>
