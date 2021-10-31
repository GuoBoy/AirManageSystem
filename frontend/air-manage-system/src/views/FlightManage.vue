<template>
	<div>
		<div class="functions">
			<el-button type="primary" @click="dialogVisible = true"
				>添加航班</el-button
			>
		</div>
		<!-- 航班列表 -->
		<el-table
			:data="
				flightList.filter(
					(data) =>
						!searchKey ||
						data.flightname
							.toLowerCase()
							.includes(searchKey.toLowerCase())
				)
			"
			style="width: 100%"
		>
			<el-table-column type="index" />
			<el-table-column
				label="航班名称"
				prop="flightname"
			/><el-table-column
				label="起飞时间"
				prop="flightstarttime"
			/><el-table-column
				label="降落时间"
				prop="flightendtime"
			/><el-table-column label="总数" prop="totalticketnum" />
			<el-table-column label="票价" prop="ticketprice" />
			<el-table-column label="折扣" prop="discountsize" />
			<el-table-column label="打折数量" prop="discountnum" />
			<el-table-column label="已售折扣票" prop="haddisnum" />
			<el-table-column label="剩余折扣票" prop="shengdiscountnum" />
			<el-table-column align="right">
				<template #header>
					<el-input
						v-model="searchKey"
						placeholder="输入航班名称进行搜索"
					/>
				</template>
				<template #default="scope">
					<el-button size="mini" @click="handleEdit(scope.row)"
						>修改</el-button
					>
					<el-button
						size="mini"
						type="danger"
						@click="handleDelete(scope.row)"
						>删除</el-button
					>
				</template>
			</el-table-column>
		</el-table>
		<!-- 添加框 -->
		<el-dialog v-model="dialogVisible" title="添加航班" width="80%">
			<el-form
				:model="newFlight"
				:rules="addRules"
				ref="addFormRef"
				label-width="100px"
			>
				<el-form-item label="航班名称" prop="flightname">
					<el-input
						v-model="newFlight.flightname"
						clearable
					></el-input>
				</el-form-item>
				<el-form-item label="航班时间">
					<el-date-picker
						v-model="timeRange"
						type="datetimerange"
						range-separator="To"
						start-placeholder="起飞时间"
						end-placeholder="降落时间"
					></el-date-picker> </el-form-item
				><el-form-item label="机票总数量">
					<el-input-number
						v-model="newFlight.totalticketnum"
						:min="1"
					/>张</el-form-item
				>
				<el-form-item label="机票价格">
					<el-input-number
						v-model="newFlight.ticketprice"
						:min="0.1"
					/>元
				</el-form-item>
				<el-form-item label="打折">
					<el-input-number
						v-model="newFlight.discountsize"
						:precision="2"
						:step="0.1"
						:max="1"
					/>折 </el-form-item
				><el-form-item label="打折数量">
					<el-input-number
						v-model="newFlight.discountnum"
						:min="1"
					/>张
				</el-form-item>
				<el-form-item>
					<el-button @click="dialogVisible = false">取消</el-button>
					<el-button type="primary" @click="onAddFlight"
						>确定</el-button
					>
				</el-form-item>
			</el-form>
		</el-dialog>
		<!-- 修改框 -->
		<el-drawer v-model="drawer" title="修改航班信息" size="60%">
			<el-form
				:model="currentFlight"
				:rules="addRules"
				label-width="100px"
				ref="addFormRef"
			>
				<el-form-item label="航班名称" prop="flightname">
					<el-input
						v-model="currentFlight.flightname"
						clearable
					></el-input>
				</el-form-item>
				<el-form-item label="航班时间">
					<el-date-picker
						v-model="currentRange"
						type="datetimerange"
						range-separator="To"
						start-placeholder="起飞时间"
						end-placeholder="降落时间"
					></el-date-picker> </el-form-item
				><el-form-item label="机票总数量">
					<el-input-number
						v-model="currentFlight.totalticketnum"
						:min="1"
					/>张</el-form-item
				>
				<el-form-item label="机票价格">
					<el-input-number
						v-model="currentFlight.ticketprice"
						:min="0.1"
					/>元
				</el-form-item>
				<el-form-item label="打折">
					<el-input-number
						v-model="currentFlight.discountsize"
						:precision="2"
						:step="0.1"
						:max="1"
					/>折 </el-form-item
				><el-form-item label="打折数量">
					<el-input-number
						v-model="currentFlight.discountnum"
						:min="1"
					/>张
				</el-form-item>
				<el-form-item>
					<el-button @click="drawer = false">取消</el-button>
					<el-button type="primary" @click="updateFlight"
						>确定</el-button
					>
				</el-form-item>
			</el-form>
		</el-drawer>
	</div>
</template>


<script>
import { onMounted, ref, computed } from "vue";
import {
	addFlight,
	getFlightList,
	updateFlightInfo,
	delFlightInfo,
} from "../api/flight";
import { ElMessage, ElMessageBox } from "element-plus";
import { useStore } from "vuex";
export default {
	setup() {
		const store = useStore();
		// 航班列表
		const flightList = ref([]);
		const searchKey = ref("");
		// 更新航班列表
		const newFlightList = () => {
			getFlightList((res) => {
				// 计算剩余折扣票
				res.data.forEach((e) => {
					var t = e.discountnum - e.haddisnum;
					e.shengdiscountnum = isNaN(t) ? e.discountnum : t;
					flightList.value.push(e);
				});
				console.log(flightList);
			});
		};
		onMounted(() => {
			newFlightList();
		});
		// 添加航班相关
		const dialogVisible = ref(false);
		const newFlight = ref({
			flightname: "",
			flightstarttime: "",
			flightendtime: "",
			totalticketnum: 1,
			discountnum: 0,
			discountsize: 1,
			ticketprice: 0,
		});
		const addRules = ref({
			flightname: [
				{
					required: true,
					message: "请输入航班名称",
					trigger: "blur",
				},
			],
		});
		const timeRange = ref([]);
		const addFormRef = ref(null);
		const onAddFlight = () => {
			addFormRef.value.validate((v) => {
				if (v) {
					var temp = timeRange.value;
					var tempData = newFlight.value;
					tempData.flightstarttime = temp[0];
					tempData.flightendtime = temp[1];
					addFlight(tempData, (res) => {
						if (res.code === 200) {
							dialogVisible.value = false;
							ElMessage.success("添加成功");
							newFlight.value = {
								flightname: "",
								flightstarttime: "",
								flightendtime: "",
								totalticketnum: 1,
								discountnum: 0,
								discountsize: 1,
								ticketprice: 0,
							};
							newFlightList();
						} else {
							ElMessage.error("请检查提交数据是否完整");
						}
					});
				} else {
					ElMessage.error("请检查提交数据是否完整");
				}
			});
		};

		// 修改航班
		const drawer = ref(false);
		const currentFlight = ref({});
		const currentRange = computed({
			get: () => {
				var temp = currentFlight.value;
				return [temp.flightstarttime, temp.flightendtime];
			},
			set: (v) => {
				currentFlight.value.flightstarttime = v[0];
				currentFlight.value.flightendtime = v[1];
			},
		});
		const handleEdit = (d) => {
			drawer.value = true;
			currentFlight.value = d;
		};
		const updateFlight = () => {
			addFormRef.value.validate((v) => {
				if (v) {
					updateFlightInfo(currentFlight.value, (res) => {
						if (res.code === 200) {
							ElMessage.success("修改成功");
							newFlightList();
							currentFlight.value = {
								flightname: "",
								flightstarttime: "",
								flightendtime: "",
								totalticketnum: 0,
								discountnum: 0,
								discountsize: 0,
								ticketprice: 0,
							};
						} else {
							ElMessage.error("修改失败");
						}
					});
				} else {
					ElMessage.error("请检查提交数据是否完整");
				}
			});
		};
		// 删除航班
		const handleDelete = (d) => {
			ElMessageBox.confirm("确定要删除航班信息吗？", "Warning", {
				confirmButtonText: "确定",
				cancelButtonText: "取消",
				type: "warning",
			})
				.then(() => {
					delFlightInfo({ ID: d.ID }, (res) => {
						if (res.code === 200) {
							ElMessage.success("删除成功");
							newFlightList();
						} else {
							ElMessage.error("删除失败");
						}
					});
				})
				.catch(() => {
					ElMessage.error("取消删除");
				});
		};

		return {
			flightList,
			searchKey,
			dialogVisible,
			newFlight,
			timeRange,
			addRules,
			addFormRef,
			onAddFlight,

			drawer,
			currentFlight,
			currentRange,
			handleEdit,
			handleDelete,
			updateFlight,
		};
	},
};
</script>

<style lang="less" scoped>
</style>