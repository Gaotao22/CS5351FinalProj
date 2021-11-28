<template>
    <div id="app" class="demo">
        <b-container fluid>
            <Navbar :dashboard="true" :logout_prop="false" />
        </b-container>
        <schedulecalendar
            :events="events"
            :dateItemRender="itemRender"
            @event-dragend="changeDate"
        ></schedulecalendar>
    </div>
</template>
<script>
import schedulecalendar from '../components/detail/scheduleCalendar/calendar'
import Navbar from '@/components/navbar/Navbar.vue'
export default {
    components: {
        schedulecalendar,
        Navbar,
    },
    data() {
        return {
            events: [
                {
                    id: 111,
                    date: '2017-04-03',
                    text: 'https://portland-my.sharepoint.com/:w:/g/personal/boweli7-c_my_cityu_edu_hk/EUgTTmBV2dlPjXUF6KYgvsAB0DqwYB0oj8kibR8ZWg7w6Q?e=emt8GU',
                    status: 1,
                },

                {
                    id: 333,
                    date: '2017-04-06',
                    text: 'https://portland-my.sharepoint.com/:w:/g/personal/boweli7-c_my_cityu_edu_hk/EVngR4CoezFPqNLeJTE6cEwB6W-AGIX0G_hWSO2UKlFXfw?e=fcwlWo',
                    status: 3,
                },
                {
                    id: 444,
                    date: '2017-04-09',
                    text: 'https://portland-my.sharepoint.com/:w:/g/personal/boweli7-c_my_cityu_edu_hk/ERXVGL3V8IRCnFR3w8FZcmcBgE4h7ZYrjM4jdH-HMFOyaA?e=TLIOVM',
                    status: 4,
                },
                {
                    id: 555,
                    date: '2017-04-11',
                    text: 'https://portland-my.sharepoint.com/:w:/g/personal/boweli7-c_my_cityu_edu_hk/EeWWeuOmXoxEhneAJtjv_Y4BCXMHR2SOHyEyc5Kxyx_RMA?e=F9TaBx',
                    status: 5,
                },
                {
                    id: 666,
                    date: '2017-04-11',
                    text: 'Sprint3',
                    status: 6,
                },
                {
                    id: 777,
                    date: '2017-04-15',
                    text: 'Meeting05-BOC Bldg 1011',
                    status: 7,
                },
                {
                    id: 888,
                    date: '2017-04-15',
                    text: 'Sprint4',
                    status: 8,
                },
                {
                    id: 999,
                    date: '2017-04-20',
                    text: 'Meeting06-BOC Bldg 1411',
                    status: 9,
                },
                {
                    id: 1010,
                    date: '2017-04-20',
                    text: 'Sprint5',
                    status: 10,
                },
            ],
            itemRender(item) {
                const h = this.$createElement
                return h('span', 'CustomRender：' + item.text)
            },
        }
    },
    methods: {
        changeDate(e, item, date) {
            const updateIndex = this.events.findIndex(
                (ele) => ele.id === item.id
            )
            this.$set(this.events, updateIndex, {
                ...this.events[updateIndex],
                date,
            })
        },
    },
    created() {
        // 显示为当月的数据
        const now = new Date()
        this.events = this.events.map((item) => {
            const d = item.date.split('-')[2]
            item.date = `${now.getFullYear()}-${now.getMonth() + 1}-${d}`
            return item
        })
    },
}
</script>
<style lang="less">
//html {
//    background: #eee;
//}
//
//body,
//html {
//    height: 100%;
//    overflow: auto;
//}
//
//body {
//    margin: 0;
//}

.demo {
    display: flex;
    flex-direction: column;
    height: 100%;
    padding: 0 30px 30px;
    box-sizing: border-box;
}

.demo-title {
    text-align: center;
}
</style>
