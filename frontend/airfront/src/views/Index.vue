<template>
	<div class="index">
		<index-header></index-header>
		<swiper></swiper>
		<!-- 查询区域 -->
		<el-card style="margin: 0 5rem">
			<el-input
				v-model="searchKEy"
				placeholder="输入航班名称进行搜索"
				clearable
			></el-input>
		</el-card>
		<!-- 航班列表显示区域 -->
		<el-card style="margin: 1rem">
			<el-table :data="flightList" style="width: 100%">
				<el-table-column type="index" /><el-table-column
					label="航班名称"
					prop="flightname"
				/>
				<el-table-column label="起飞时间" prop="flightstarttime" />
				<el-table-column label="降落时间" prop="flightendtime" />
				<el-table-column label="票价" prop="ticketprice" />
				<el-table-column label="载客量" prop="totalticketnum" />
				<el-table-column label="打折" prop="discountsize" />
				<el-table-column align="right">
					<template #default="scope">
						<el-button
							size="mini"
							type="primary"
							@click="onPurchase(scope.row)"
							>折扣购买</el-button
						>
					</template>
				</el-table-column>
			</el-table>
		</el-card>
		<!-- 支付 -->
		<el-dialog v-model="dialogVisible" title="请微信扫码支付" width="80%">
			<div class="form">
				<div class="form-tags">
					<el-tag type="warning">提示，折扣票不可退改签</el-tag>
					<el-tag type="danger"
						>请在24h之内支付，否则将放入票库</el-tag
					>
				</div>
				<div class="qrcode">
					<img src="../assets/zhifu.png" />
				</div>
				<el-button type="primary" class="btn" @click="payed"
					>已支付</el-button
				>
			</div>
		</el-dialog>
	</div>
</template>

<script>
import { computed, ref, onMounted } from "vue";
import IndexHeader from "../components/IndexHeader.vue";
import Swiper from "../components/Swiper.vue";
import { getFlightList } from "../api/flight";
import { addTicket } from "../api/ticket";
import { ElMessage } from "element-plus";
import { useStore } from "vuex";
export default {
	components: {
		IndexHeader,
		Swiper,
	},
	setup() {
		const store = useStore();
		const tempFlights = ref([]);
		const newFlightList = () => {
			getFlightList((res) => {
				if (res.code === 200) {
					tempFlights.value = res.data;
					ElMessage.success("初始化航班列表成功");
				} else {
					ElMessage.error("初始化航班列表失败，请重试");
				}
			});
		};
		onMounted(() => {
			newFlightList();
		});
		// 搜索
		const searchKEy = ref("");
		const flightList = computed(() =>
			tempFlights.value.filter((e) =>
				e.flightname
					.toLowerCase()
					.includes(searchKEy.value.toLowerCase())
			)
		);
		// 购买
		const dialogVisible = ref(false);
		const currentOrder = ref({});
		const onPurchase = (f) => {
			dialogVisible.value = true;
			currentOrder.value = f;
		};
		const payed = () => {
			dialogVisible.value = false;
			addTicket(
				{
					flightfid: currentOrder.value.ID,
					useridcard: store.getters.getToken,
				},
				(res) => {
					if (res.code === 200) {
						ElMessage.success("购买成功");
					} else {
						ElMessage.error("购买失败");
					}
				}
			);
		};

		return {
			flightList,
			searchKEy,
			dialogVisible,
			onPurchase,
			payed,
		};
	},
};
</script>

<style lang="less" scoped>
.form {
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	width: 100%;
	height: 100%;
	.el-tag {
		margin-right: 1rem;
	}
	.qrcode {
		width: 30%;
		height: 0;
		padding-bottom: 30%;
		position: relative;
		img {
			width: 100%;
			height: 100%;
			position: absolute;
		}
	}
}
</style>