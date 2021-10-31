<template>
	<el-table
		:data="
			userlist.filter(
				(data) =>
					!search ||
					data.realname.toLowerCase().includes(search.toLowerCase())
			)
		"
		style="width: 100%"
	>
		<el-table-column type="index" />
		<el-table-column label="用户名" prop="username" />
		<el-table-column label="真实姓名" prop="realname" />
		<el-table-column label="电话" prop="phone" />
		<el-table-column label="密码" prop="password" />
		<el-table-column label="用户类型" prop="issuper" />
		<el-table-column align="right">
			<template #header>
				<el-input
					v-model="search"
					size="mini"
					placeholder="输入用户真实姓名进行搜索"
				/>
			</template>
			<template #default="scope">
				<el-button
					size="mini"
					type="danger"
					@click="handleDelete(scope.row)"
					>删除</el-button
				>
			</template>
		</el-table-column>
	</el-table>
</template>


<script>
import { onMounted, ref } from "vue";
import { getUsers, delUser } from "../api/user";
import { ElMessage, ElMessageBox } from "element-plus";
import { useStore } from "vuex";
export default {
	setup() {
		const store = useStore();
		// 航班列表
		const userlist = ref([]);
		// 搜索关键词
		const search = ref("");

		const newUserList = () => {
			getUsers((res) => {
				userlist.value = res.data;
			});
		};

		const handleDelete = (d) => {
			ElMessageBox.confirm("确定要删除信息吗？", "Warning", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				type: "warning",
			})
				.then(() => {
					delUser({ ID: d.ID }, (res) => {
						if (res.code === 200) {
							ElMessage({
								type: "success",
								message: "删除成功",
							});
							newUserList();
						} else {
							ElMessage({
								type: "info",
								message: "删除失败",
							});
						}
					});
				})
				.catch(() => {
					ElMessage({
						type: "info",
						message: "取消删除",
					});
				});
		};
		// 渲染
		onMounted(() => {
			newUserList();
		});

		return {
			userlist,
			search,
			handleDelete,
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