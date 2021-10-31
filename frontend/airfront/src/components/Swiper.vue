<template>
	<!-- swpier区域 -->
	<el-carousel height="300px" type="card">
		<el-carousel-item v-for="(img, index) in swiperImgs" :key="index">
			<!-- content -->
			<el-image :src="img"></el-image>
		</el-carousel-item>
	</el-carousel>
</template>

<script>
import { ref } from "vue";
export default {
	setup() {
		// 轮播图
		const swiperImgs = ref([
			require("../assets/bg.jpg"),
			require("../assets/bg.jpg"),
			require("../assets/bg.jpg"),
			require("../assets/bg.jpg"),
			require("../assets/bg.jpg"),
			require("../assets/bg.jpg"),
		]);

		return {
			swiperImgs,
		};
	},
};
</script>