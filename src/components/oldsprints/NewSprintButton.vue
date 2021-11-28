<template>
    <div class="text-font">
        <b-jumbotron v-if="numOfSprints === 0">
            <h1>No Sprints in this Project</h1>
            <div
                v-if="
                    projectOverview.project.isDone === false &&
                    teamRole() != 'Developer'
                "
            >
                <b>Start a new sprint!</b><br /><br />
                <b-button
                    size="lg"
                    variant="success"
                    class="m-md-3"
                    :to="{
                        name: 'NewSprint',
                        params: { id: $route.params.id },
                    }"
                    >Create New Sprint</b-button
                >
            </div>
        </b-jumbotron>
        <b-row v-if="numOfSprints > 0">
            <b-col class="text-right">
                <el-button
                    @click="drawer = true"
                    type="primary"
                    size="large"
                    style="margin-left: 16px;"
                >
                    Burndown Chart
                </el-button>
                <el-drawer
                    v-if="drawer"
                    :visible.sync="drawer"
                    :with-header="false"
                >
                    <dv-charts
                        :option="initburnchart()"
                        style="width: 100%; height: 100%;"
                    />
                </el-drawer>
                <b-button
                    v-if="
                        projectOverview.project.isDone === false &&
                        teamRole() != 'Developer'
                    "
                    size="lg"
                    variant="success"
                    class="m-md-3"
                    :to="{
                        name: 'NewSprint',
                        params: { id: $route.params.id },
                    }"
                >
                    Create New Sprint
                </b-button>
            </b-col>
        </b-row>
    </div>
</template>

<script>
import axios from 'axios'
import echarts from 'echarts'
import Vue from 'vue'
Vue.prototype.$echarts = echarts

export default {
    name: 'NewSprintButton',
    props: {
        numOfSprints: Number,
    },
    data() {
        return {
            drawer: false,
            projectOverview: {
                project: {
                    id: 0,
                    isDone: false,
                },
            },
            option1: {},
        }
    },
    methods: {
        getSprintInfo() {
            const self = this
            axios
                .get(
                    this.$url +
                        '/users/' +
                        localStorage.getItem('userId') +
                        '/projects/' +
                        this.$route.params.id,
                    {
                        headers: {
                            Authorization:
                                'Bearer ' + localStorage.getItem('auth_token'),
                            'Content-Type': 'application/json',
                        },
                    }
                )
                .then(function (response) {
                    if (response.data.serverErrorMessage) {
                        console.log(response.data.serverErrorMessage)
                    } else {
                        self.projectOverview = response.data
                    }
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        teamRole() {
            return localStorage.getItem('teamRole')
        },
        async getburnchart() {
            let xaxis = []
            let yaxis = []
            let yaxis_length = 0
            await axios
                .get(
                    this.$url +
                        '/users/' +
                        localStorage.getItem('userId') +
                        '/projects/' +
                        this.$route.params.id +
                        '/remaintime',
                    {
                        headers: {
                            Authorization:
                                'Bearer ' + localStorage.getItem('auth_token'),
                            'Content-Type': 'application/json',
                        },
                    }
                )
                .then(function (response) {
                    if (response.data.editorErrorMessage) {
                        console.log(response.data.serverErrorMessage)
                    } else {
                        // eslint-disable-next-line no-empty
                        xaxis = new Array(response.data.length)
                        yaxis = new Array(response.data.length)
                        for (var i = 0; i < response.data.length; i++) {
                            xaxis[i] = 'sprint' + String(i+1)
                        }
                        console.log('你给我的燃尽图time', response.data)
                        response.data.forEach(function (element) {
                            console.log(element)
                            if (yaxis_length > 0) {
                                yaxis[yaxis_length] =
                                    yaxis[yaxis_length - 1] - element
                            } else {
                                yaxis[yaxis_length] = 470 - element
                            }
                            yaxis_length = yaxis_length + 1
                        })
                        console.log(xaxis)
                        console.log(yaxis)
                    }
                })
                .catch(function (error) {
                    console.log(error)
                })
            return { x: xaxis, y: yaxis }
        },
        initburnchart() {
            console.log('1')
            this.getburnchart().then((x) => {
                console.log('2')
                this.option1 = {
                    title: {
                        text: 'Burntdown Chart',
                    },
                    xAxis: {
                        name: 'sprint',
                        data: x['x'],
                    },
                    yAxis: {
                        name: 'Left time',
                        data: 'value',
                    },
                    series: [
                        {
                            data: x['y'],
                            type: 'line',
                        },
                    ],
                }
                console.log('3')
            })
            return this.option1
        },
    },
    mounted() {
        this.getSprintInfo()
    },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Quicksand');

.text-font {
    font-family: Quicksand;
}
</style>
