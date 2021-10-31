<template>
	<div>
		<index-header></index-header>
		<el-card style="margin: 1rem">
			<el-table border stripe :data="cartList" fit>
				<el-table-column type="index"> </el-table-column>
				<el-table-column label="机票编号" prop="ID"> </el-table-column>
				<el-table-column label="航班ID" prop="flightfid">
				</el-table-column>
				<el-table-column align="right">
					<template #default="scope">
						<el-button
							size="mini"
							type="danger"
							@click="onDeleteTicket(scope.row)"
							>删除</el-button
						>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
	</div>
</template>

<script>
import { onMounted, ref } from "vue";
import IndexHeader from "../components/IndexHeader.vue";
import { getTicketList, delTicket } from "../api/ticket";
import { ElMessage, ElMessageBox } from "element-plus";
import { useStore } from "vuex";
export default {
	components: {
		IndexHeader,
	},
	setup() {
		const store = useStore();
		// 购物车列表
		const cartList = ref([]);

		// 更新机票列表
		const newTicketList = () => {
			getTicketList({ token: store.getters.getToken }, (res) => {
				cartList.value = res.data;
				console.log(res.data);
			});
		};

		// 删除机票
		const onDeleteTicket = (d) => {
			ElMessageBox.confirm("确定要退票吗？", "Warning", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				type: "warning",
			})
				.then(() => {
					delTicket({ ID: d.ID }, (res) => {
						if (res.code === 200) {
							ElMessage({
								type: "success",
								message: "退票成功",
							});
							newTicketList();
						} else {
							ElMessage({
								type: "info",
								message: "退票失败",
							});
						}
					});
				})
				.catch(() => {
					ElMessage({
						type: "info",
						message: "取消退票",
					});
				});
		};

		onMounted(() => {
			newTicketList();
		});

		return {
			cartList,
			onDeleteTicket,
		};
	},
};
</script>