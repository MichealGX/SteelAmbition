<template>
	<view style="padding: 10px;">
		<view v-for="(item, index) in displayedList" :key="index" style="width: 100%; height: 80px; display: flex; align-items: center; justify-content: space-around; margin-top: 15px; background: #e9e9e9; border-radius: 10px;">
			<span>{{ item.id }}</span>
			<span>{{ item.name }}</span>
		</view>
		<view style="width: 100%; height: 80px; display: flex; align-items: center; justify-content: space-around">
			<button @click="prevPage" :disabled="currentPage === 1">上一页</button>
			<span>{{ currentPage }}页</span>
			<button @click="nextPage" :disabled="currentPage === totalPages">下一页</button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				listVm: [
					{ id: '1', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '2', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '3', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '4', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '5', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '6', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '7', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '8', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '9', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '10', name: '嗨(｡･∀･)ﾉﾞ嗨' },
					{ id: '11', name: '嗨(｡･∀･)ﾉﾞ嗨' },
				],
				currentPage: 1,
				pageSize: 5,
			}
		},
		computed: {
			totalPages() {
				return Math.ceil(this.listVm.length / this.pageSize);
			},
			displayedList() {
				const start = (this.currentPage - 1) * this.pageSize;
				const end = start + this.pageSize;
				return this.listVm.slice(start, end);
			}
		},
		methods: {
			prevPage() {
				if (this.currentPage > 1) {
					this.currentPage--;
				}
			},
			nextPage() {
				if (this.currentPage < this.totalPages) {
					this.currentPage++;
				}
			}
		}
	}
</script>

<style>
</style>