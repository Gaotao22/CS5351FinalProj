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
                <!--                <b-button-->
                <!--                    v-if="-->
                <!--                        projectOverview.project.isDone === false &&-->
                <!--                        teamRole() !== 'Developer'-->
                <!--                    "-->
                <!--                    size="lg"-->
                <!--                    variant="success"-->
                <!--                    class="m-md-3"-->
                <!--                    v-b-modal.modal-1-->
                <!--                >-->
                <!--                    Burndown Chart-->
                <!--                </b-button>-->
                <!--                &lt;!&ndash;                <burntchart&ndash;&gt;-->
                <!--                &lt;!&ndash;                    v-on:&ndash;&gt;-->
                <!--                &lt;!&ndash;                    :modalId = "'burnt_chart'"&ndash;&gt;-->
                <!--                &lt;!&ndash;                />&ndash;&gt;-->
                <!--                <b-modal id="modal-1" title="BootstrapVue">-->
                <!--                    <div style="height: 300px; width: 450px;">-->
                <!--                        <dv-charts-->
                <!--                            :option="option1"-->
                <!--                            style="width: 100%; height: 100%;"-->
                <!--                        />-->
                <!--                    </div>-->
                <!--                </b-modal>-->
                <el-button
                    @click="drawer = true"
                    type="primary"
                    size="large"
                    style="margin-left: 16px;"
                >
                    Burndown Chart
                </el-button>
                <el-drawer :visible.sync="drawer" :with-header="false">
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
            option1: {
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
            },
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
        initburnchart() {
            axios
                .get(
                    this.$url +
                        '/users/' +
                        localStorage.getItem('userId') +
                        '/projects/' +
                        this.$route.params.id +
                        'remaintime',
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
                        console.log('你给我的燃尽图time', response.data)
                        response.data.forEach(function(element) {
                            console.log(element)
                        })
                    }
                })
                .catch(function (error) {
                    console.log(error)
                })
            return {
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
