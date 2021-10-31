<template>
	<router-view></router-view>
</template>

<script>
import { onMounted } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
export default {
	setup() {
		const store = useStore();
		const router = useRouter();
		// 检查登录
		const checkLogin = () => {
			if (
				store.getters.getUserName !== "" &&
				store.getters.getToken !== ""
			)
				router.push("index");
		};
		onMounted(() => {
			checkLogin();
		});
	},
};
</script>

