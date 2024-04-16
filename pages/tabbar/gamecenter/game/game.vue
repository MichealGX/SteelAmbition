<template>
  <view>
 
      <view class="video-container">
		  
      </view>

	<view class="container">
	  <image class="button left-button" src="/static/left-arrow.png" @tap="sendCommand('left')" />
	  <image class="button right-button" src="/static/right-arrow.png" @tap="sendCommand('right')" />
	  <image class="button up-button" src="/static/up-arrow.png" @tap="sendCommand('forward')" />
	  <image class="button down-button" src="/static/down-arrow.png" @tap="sendCommand('backward')" />
	  <button class="fire-mode-button" @tap="navigateToFireMode">切换开火</button>
	  </view>
	  <!-- 显示车辆状态信息 -->
	  <view class="status-container">
	    <text>车辆状态: {{ vehicleStatus.status }}</text>
	    <text>核心能量: {{ vehicleStatus.core_e }}</text>
	    <text>武器能量: {{ vehicleStatus.weapon_e }}</text>
	    <text>防御能量: {{ vehicleStatus.defence_e }}</text>
	    <text>行走能量: {{ vehicleStatus.walking_e }}</text>
	  </view>
	  
	</view>
	
</template>

<script>
const udpClient = uni.requireNativePlugin('udp-client');
export default {
  data() {
    return {
      socketPort: 5173,
      vehicleStatus: {} // 存储车辆状态信息
    };
  },
  mounted() {
    // 页面加载时初始化车辆状态
    this.initVehicleStatus();
  },
  onLoad() {
    /**
     * 如果需要使用 hexString， 请在 init 之前调用 
     * 之后会在返回数据多一个 hex 字段
     */
    udpClient.setUseHex(true);

    /**
     * 设置接受的字节大小，单位是 byte，默认 1024
     * 请按需设置，过大可能会影响性能，或导致奔溃
     */
    udpClient.setByteSize(2048);

    /**
     * 在设备初始化，监听 10005 端口。
     * 假设本设备 IP 为 192.168.2.35
     * 那么服务端，或者其设备，就可以给 192.168.2.35:10005 发送消息了
     */
    udpClient.init(this.socketPort, this.onSocketMsg, this.onSocketError);
  },
  methods: {
    onSocketMsg(resData) {
      // resData 的数据结构：{ host, port, data, hex }
      console.log("接收到消息: " + resData);
      // 收到消息如果想响应
      udpClient.send({
        host: resData.host,
        port: resData.port,
        data: bytebuffer({contype: 1, msg: "我收到消息啦～"}),
      });
    },
    onSocketError(errMsg) {
      console.error("socket 异常：" + errMsg);
    },
    sendCommand(direction) {
      // 发送操作数据给小车服务器
      uni.request({
        url: 'https://127.0.0.1:8080/platform/vehicles/your_vehicle_id/operation/move', // 替换为实际的车辆ID
        method: 'POST',
        data: {
          type: 0,
          direction: direction
        },
        success(res) {
          console.log('操作成功:', res.data);
        },
        fail(err) {
          console.error('操作失败:', err);
        }
      });
    },
    navigateToFireMode() {
      // 跳转到开火模式页面
      uni.navigateTo({
        url: '/pages/tabbar/gamecenter/game/fire/fire'
      });
    },
    initVehicleStatus() {
      // 初始化车辆状态
      uni.request({
        url: 'https://127.0.0.1:8080/platform/vehicles/your_vehicle_id/operation/status', // 替换为实际的车辆ID
        method: 'GET',
        success: (res) => {
          console.log('车辆状态信息:', res.data);
          this.vehicleStatus = res.data; // 更新为直接存储服务器返回的状态信息
        },
        fail: (err) => {
          console.error('获取车辆状态信息失败:', err);
        }
      });
    }
  }
};
</script>



<style>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  transform: translateY(-200px); 
}

.button {
  width: 50px;
  height: 50px;
  margin: 10px;
}


.right-button {
  position: relative;
  top: 75px;
  right: -55px;
}

.left-button {
  position: relative;
  top: 145px;
  right: 55px;
}
.up-button {
  position: relative;
  top: -40px;
}

.down-button {
  position: relative;
  top: -20px;
}
.fire-mode-button {
  width: 150px;
  height: 50px;
  margin-top: 20px;
}

.status-container {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-top: 20px;
}
.video-container {
  border-radius: 8px; /* 设置矩形方框的圆角 */
  padding: 16px; /* 设置内边距 */
  height: 200px; /* 设置高度 */
  border: 2px solid #000000; /* 添加边框 */
}
</style>