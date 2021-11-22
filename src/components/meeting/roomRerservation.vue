<template>
    <div class="resRoomMain">
        <div class="roomReserveMain">
            <v-room-con-filter
                @change="
                    getRoomLst()
                    getPagination()
                "
            ></v-room-con-filter>
            <v-room-plugin
                @click="getRoomLst()"
                :room-lst="roomLst"
            ></v-room-plugin>

            <div id="roomResPag">
                <el-pagination
                    @current-change="getRoomLst()"
                    :current-page="pagination.currentPage"
                    layout="total,  prev, pager, next, jumper"
                    :total="pagination.total"
                >
                </el-pagination>
            </div>
            <div id="resRoomInfo">
                <h1>Booking Information</h1>
                <ul>
                    <li class="resRoomInfoItem">
                        <span class="resRoomInfoItemTitle">Room Area：</span>
                        {{ resRoomInfo.region }}
                    </li>
                    <li class="resRoomInfoItem">
                        <span class="resRoomInfoItemTitle"
                            >Meeting Room Name：</span
                        >
                        {{ resRoomInfo.roomName }}
                    </li>
                    <li class="resRoomInfoItem">
                        <span class="resRoomInfoItemTitle">Booking Time：</span>
                        {{ resRoomInfo.date }}
                        <span v-if="resRoomInfo.week">（</span>
                        {{ resRoomInfo.week }}
                        <span v-if="resRoomInfo.week">）</span>
                        {{ resRoomInfo.startTime }}
                        <span v-if="resRoomInfo.startTime">——</span>
                        {{ resRoomInfo.endTime }}
                    </li>
                    <li class="resRoomInfoItem">
                        <span class="resRoomInfoItemTitle">User Name：</span>
                        {{ resRoomInfo.user }}
                    </li>
                </ul>
            </div>
            <div id="resBtnWrap">
                <el-button type="primary">BOOKING</el-button>
            </div>
        </div>
    </div>
</template>

<script>
import vRoomConFilter from '@/components/meeting/vRoomConFilter'
import vRoomPlugin from '@/components/meeting/vRoomPlugin'
import getRoomLst from '../../mock/room.js'
import getUserName from '../../mock/getUserName.js'
import getPagination from '../../mock/pagination.js'

import Vue from 'vue'
import ElementUI from 'element-ui'
import 'font-awesome/css/font-awesome.css'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(ElementUI)
export default {
    name: 'roomResPlug',
    data() {
        return {
            regionName: null,
            roomLst: [],
            resRoomInfo: {
                region: null,
                date: null,
                week: null,
                startTime: null,
                endTime: null,
                roomName: null,
                user: null,
            },
            userName: null,
            pagination: {
                total: 0,
                currentPage: 0,
            },
        }
    },
    components: {
        vRoomConFilter,
        vRoomPlugin,
    },
    methods: {
        getRoomLst() {
            this.roomLst = getRoomLst().roomLst //传递会议室信息
        },
        getUserName() {
            this.userName = getUserName().userName
        },
        getPagination() {
            let pagination = getPagination()
            this.pagination.total = pagination.total
            this.pagination.currentPage = pagination.currentPage
        },
    },
    mounted() {
        this.getRoomLst() //获取会议室信息
        this.getUserName()
        this.getPagination()
    },
    created() {},
}
</script>

<style scoped lang="less">
@sumW: 900px;
@roomTopSideW: 120px;
@topRightW: @sumW - @roomTopSideW;
@fontColor:#606266;
@centerW:1000px;
.resRoomMain {
    width: @centerW;
    margin: 20px auto;
    background-color: #fff;
    padding: 20px 0;
}

.roomReserveMain {
    width: @sumW;
    min-width: @sumW;
    margin: 0 auto;
}

//预订信息卡片
#resRoomInfo {
    margin-top: 30px;
    margin-bottom: 30px;
    color: @fontColor;

    h1 {
        font-size: 16px;
        margin-bottom: 10px;
    }

    .resRoomInfoItem {
        padding: 5px 0;

        .resRoomInfoItemTitle {
            margin-right: 10px;
        }
    }
}

#roomResPag {
    display: flex;
    justify-content: flex-end;
    padding: 15px 0;
}

#resBtnWrap {
    text-align: center;
    padding: 20px 0;
}

// screen >= 1200
@media screen and (min-width: 1200px) {
}

// 1200>= screen >=992
@media screen and (max-width: 1200px) {
}

@media screen and (max-width: 992px) {
}

@media screen and (max-width: 768px) {
}

@media screen and (max-width: 480px) {
}
</style>
