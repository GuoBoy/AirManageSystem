<template>
	<el-table
		:data="
			orderlist.filter(
				(data) =>
					!search ||
					data.username.toLowerCase().includes(search.toLowerCase())
			)
		"
		style="width: 100%"
	>
		<el-table-column type="index" />
		<el-table-column label="订单ID" prop="tid" />
		<el-table-column label="航班名称" prop="flightname" />
		<el-table-column label="用户名" prop="username" />
		<el-table-column label="价格" prop="price" />
		<el-table-column label="预定时间" prop="UpdatedAt" />
		<el-table-column align="right">
			<template #header>
				<el-input
					v-model="search"
					size="mini"
					placeholder="输入用户姓名进行搜索"
				/>
			</template>
			<el-button size="mini" type="success" @click="sendNotify"
				>发送提醒</el-button
			>
		</el-table-column>
	</el-table>
</template>


<script>
import { onMounted, ref } from "vue";
import { getTicketList } from "../api/ticket";
import { ElMessage, ElMessageBox } from "element-plus";
import { useStore } from "vuex";
export default {
	setup() {
		// 订单列表
		const orderlist = ref([]);
		// 搜索关键词
		const search = ref("");

		const newOrderList = () => {
			getTicketList((res) => {
				if (res.code === 200) {
					ElMessage.success("获取机票数据成功");
					console.log(res.data);
					orderlist.value = res.data;
				} else {
					ElMessage.error("获取机票列表失败");
				}
			});
		};

		const sendNotify = () => {
			ElMessageBox.confirm("确定要发送短信提醒吗？", "Warning", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				type: "warning",
			})
				.then(() => {
					setTimeout(() => {
						ElMessage.success("发送成功");
					}, 1500);
				})
				.catch(() => {
					ElMessage({
						type: "info",
						message: "取消发送",
					});
				});
		};
		// 渲染
		onMounted(() => {
			newOrderList();
		});

		return {
			orderlist,
			search,
			sendNotify,
		};
	},
};
</script>

<style lang="less" scoped>
.dialog-form {
	.dialog-item {
		margin: 1rem 0;
		width: 100%;
		.item {
			width: 50%;
		}
	}
}
</style>