<template>
	<div class="header">
		<div class="logo"><img src="../assets/bg.jpg" />理工航空欢迎您</div>
		<router-link to="login" v-if="username === ''">
			<el-button type="primary">登录/注册</el-button>
		</router-link>
		<el-dropdown @command="handleCommand" v-else>
			<span class="el-dropdown-link">
				欢迎 {{ username }}<i class="el-icon-arrow-down"></i>
			</span>
			<template #dropdown>
				<el-dropdown-menu>
					<el-dropdown-item command="center">首页</el-dropdown-item>
					<el-dropdown-item command="cart">购物车</el-dropdown-item>
					<el-dropdown-item divided command="exit"
						>退出登录</el-dropdown-item
					>
				</el-dropdown-menu>
			</template>
		</el-dropdown>
	</div>
</template>

<script>
import { computed, toRefs } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

export default {
	setup(props, context) {
		const store = useStore();
		const router = useRouter();

		const username = computed(() => {
			try {
				return store.getters.getUserName;
			} catch (error) {
				return "";
			}
		});

		const handleCommand = (e) => {
			switch (e) {
				case "center":
					router.push("index");
					break;
				case "cart":
					router.push("cart");
					break;
				case "exit":
					console.log("退出登录");
					store.commit("updateUserInfo", { u: "", t: "" });
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