<template>
	<el-container class="backmanage">
		<el-header>
			<back-manage-header></back-manage-header>
		</el-header>
		<el-container>
			<el-aside width="100px">
				<el-menu
					class="index-menu"
					:default-active="activeMenu"
					router
					v-for="(menu, index) in menus"
					:key="index"
					@select="onSelect"
				>
					<!-- 菜单栏项目遍历 -->
					<el-menu-item :index="menu.path">
						<template #title>
							{{ menu.title }}
						</template>
					</el-menu-item>
				</el-menu></el-aside
			>
			<el-main>
				<!-- Main content -->
				<el-card>
					<router-view></router-view>
				</el-card>
			</el-main>
		</el-container>
	</el-container>
</template>


<script>
import { ref } from "vue";
import BackManageHeader from "../components/BackManageHeader.vue";
export default {
	components: {
		BackManageHeader,
	},
	setup() {
		// 菜单
		const menus = ref([
			{
				title: "航班管理",
				path: "/index/flightmanage",
			},
			{
				title: "用户管理",
				path: "/index/usermanage",
			},
			{
				title: "订单管理",
				path: "/index/ordermanage",
			},
		]);
		const activeMenu = ref(menus.value[0].path);
		const onSelect = (s) => (activeMenu.value = s);
		return {
			menus,
			activeMenu,
			onSelect,
		};
	},
};
</script>

<style lang="less" scoped>
.el-card {
	margin: 0 1rem;
	min-height: 42rem;
}
.el-header {
	padding: 0;
}
</style>
