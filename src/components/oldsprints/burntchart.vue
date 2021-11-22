<template>
    <b-modal :id="modal-1">
        <div style="height: 300px; width: 450px;">
            <dv-charts :option="option" style="width: 100%; height: 100%;" />
        </div>
    </b-modal>
</template>

<script>
export default {
    name: 'BurnChart',
    props: {
        modalId: String,
    },
    data() {
        return {
            appear: 1,
            option: {},
        }
    },
    methods: {
        initburnchart() {
            this.option = {
                title: {
                    text: 'Burntdown Chart',
                },
                xAxis: {
                    name: 'sprint',
                    data: [
                        'sprint1',
                        'sprint2',
                        'sprint3',
                        'sprint4',
                        'sprint5',
                        'sprint6',
                        'sprint7',
                    ],
                },
                yAxis: {
                    name: 'Left time',
                    data: 'value',
                },
                series: [
                    {
                        data: [350, 320, 300, 250, 200, 180, 160],
                        type: 'line',
                    },
                ],
            }
        },
    },
    created() {
      console.log('xuanran')
      this.initburnchart()

    },
}
</script>

<style scoped></style>
