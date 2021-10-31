<template>
	<div class="header">
		<div class="logo"><img src="../assets/bg.jpg" />航空管理系统后台</div>
		<el-dropdown @command="handleCommand">
			<span class="el-dropdown-link">
				欢迎 {{ username }}<i class="el-icon-arrow-down"></i>
			</span>
			<template #dropdown>
				<el-dropdown-menu>
					<el-dropdown-item command="center"
						>个人中心</el-dropdown-item
					>
					<el-dropdown-item divided command="exit"
						>退出登录</el-dropdown-item
					>
				</el-dropdown-menu>
			</template>
		</el-dropdown>
	</div>
</template>

<script>
import { toRefs } from "vue";
import { useStore } from "vuex";

export default {
	setup() {
		const store = useStore();

		const username = store.getters.getUserName;

		const handleCommand = (e) => {
			switch (e) {
				case "center":
					alert("跳转到个人中心");
					break;
				case "exit":
					store.commit("updateUserInfo", { u: "", t: "" });
					console.log(store.getters.getUserName);
					window.location.href = "/";
					break;
			}
		};

		return {
			username,
			handleCommand,
		};
	},
};
</script>

<style lang="less" scoped>
.header {
	height: 3rem;
	width: 100%;
	display: flex;
	justify-content: space-between;
	align-items: center;
	background-color: rgb(135, 210, 253);
	.logo {
		width: 3rem;
		height: 3rem;
		display: flex;
		white-space: nowrap;
		align-items: center;
		img {
			width: 100%;
			height: 100%;
			border-radius: 1.5rem;
			margin-right: 1rem;
		}
	}
}
</style>