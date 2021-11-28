<template>
    <div>
        <header-bar
            :year="year"
            :month="month"
            @updateValue="updateView"
      ></header-bar>
        <div class="schedule-calendar">
            <week :startWeek="startWeek"></week>
            <month
                :year="year"
                :month="month"
                :startWeek="startWeek"
                :data="events"
                :itemRender="dateItemRender"
            ></month>
        </div>
    </div>
</template>
<script>
import { EventBus } from '@/components/detail/scheduleCalendar/utils'
import headerBar from '@/components/detail/scheduleCalendar/header'
import week from '@/components/detail/scheduleCalendar/week'
import month from '@/components/detail/scheduleCalendar/month'

export default {
    name: 'schedulecalendar',
    components: {
        headerBar,
        week,
        month,
    },
    props: {
        startMonth: String,
        startWeek: {
            type: Number,
            default: 1,
        },
        events: Array,
        dateItemRender: Function,
    },
    data() {
        return {
            year: new Date().getFullYear(),
            month: new Date().getMonth(),
            dragItem: null,
        }
    },
    methods: {
        updateView(year, month) {
            this.year = year
            this.month = month
        },
        // eslint-disable-next-line no-unused-vars
        cellDragenter(e, date, type, index) {
            this.$emit('event-dragenter', e, this.dragItem, date)
        },
        // eslint-disable-next-line no-unused-vars
        itemDragstart(e, item, date, type) {
            this.dragItem = item
            this.$emit('event-dragstart', e, item, date)
        },
        // eslint-disable-next-line no-unused-vars
        itemDrop(e, date, type, index) {
            if (!this.dragItem) return
            console.log('[event-dragend]:', this.dragItem, date)
            this.$emit('event-dragend', e, this.dragItem, date)
        },
        itemClick(e, item) {
            console.log('[event-click]:', item)
            this.$emit('event-click', e, item)
            window.open(item.text,'_blank')
        },
        dateClick(e, date) {
            console.log('[date-click]:', date)
            this.$emit('date-click', e, date)
        },
    },
    created() {
        EventBus.$on('cell-dragenter', this.cellDragenter)
        EventBus.$on('item-dragstart', this.itemDragstart)
        EventBus.$on('item-drop', this.itemDrop)
        EventBus.$on('item-click', this.itemClick)
        EventBus.$on('date-click', this.dateClick)
    },
    destoryed() {
        EventBus.$off()
    },
}
</script>
<style lang="less">
@import './src/components/detail/scheduleCalendar/variables.less';

.schedule- {
    &calendar {
        position: relative;
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        min-width: @sc-cell-min-width * 7;
        color: @sc-base-color;
        font-size: @sc-base-font-size;
        box-shadow: @sc-box-shadow;

        *,
        *::before,
        *::after {
            box-sizing: border-box;
        }

        button {
            border: 0;
            outline: none;
            cursor: pointer;
            background: transparent;
        }

        &body {
            position: relative;
            flex: 1;
            top: 50px;
            width: 100%;
            overflow: hidden;
            background: @sc-body-color;
        }
    }
}
</style>
