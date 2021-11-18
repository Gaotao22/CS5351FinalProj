<template>
    <b-col class="text-right">
        <b-button @click="openModal()">
            <img src="@/assets/edit.png" class="btn-img" />
            Edit
        </b-button>
        <b-modal v-model="modalShow" title="Edit" @ok="updateTask()">
            <b-form class="text-left" id="updateEpic">
                <b-form-group
                    label="Description:"
                    :label-for="'editTask' + Task_id"
                >
                    <b-form-input
                        :id="'editTask' + Task_id"
                        type="text"
                        v-model="updateTask_desc"
                        required
                    >
                    </b-form-input>
                </b-form-group>
                <b-form-group
                    label="complete time:"
                    :label-for="'editTask' + Task_id"
                >
                    <b-form-input
                        :id="'editTask' + Task_id"
                        type="number"
                        v-model="updateTask_comtime"
                        required
                    >
                    </b-form-input>
                </b-form-group>
            </b-form>
            <template v-slot:modal-footer="{ remove, cancel, ok }">
                <b-btn class="mr-auto" variant="danger" @click="deleteTask()">
                    Delete Task (and its Issues)</b-btn
                >
                <b-btn variant="secondary" @click="cancel()">Cancel</b-btn>
                <b-btn variant="primary" @click="ok()">OK</b-btn>
            </template>
        </b-modal>
    </b-col>
</template>

<script>
import axios from 'axios'

export default {
    name: 'EditTask',
    props: {
        desc: String,
        PBI_id: Number,
        state: Number,
        Task_id: Number,
        comtime: Number,
    },
    data: function () {
        return {
            updateTask_desc: '',
            modalShow: false,
            updateTask_comtime: '',
        }
    },
    methods: {
        openModal() {
            //console.log(typeof(this.epicId));
            this.updateTask_desc = this.desc
            this.modalShow = !this.modalShow
            this.updateTask_comtime = this.comtime
        },
        updateTask() {
            const self = this
            let data = {
                storyId: this.PBI_id,
                description: this.updateTask_desc,
                state: this.state,
                id: this.Task_id,
                comtime: this.updateTask_comtime,
            }
            axios
                .put(
                    this.$url +
                        '/users/' +
                        localStorage.getItem('userId') +
                        '/projects/' +
                        this.$route.params.id +
                        '/pbis/' +
                        this.PBI_id +
                        '/tasks',
                    data,
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
                        console.log('我给你的edittask',data)
                        console.log('你给我的edittask',response.data)
                        self.$emit('editTask', response.data)
                    }
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
        deleteTask() {
            const self = this
            axios
                .delete(
                    this.$url +
                        '/users/' +
                        localStorage.getItem('userId') +
                        '/projects/' +
                        this.$route.params.id +
                        '/pbis/' +
                        this.PBI_id +
                        '/tasks/' +
                        this.Task_id,
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
                        self.modalShow = false
                        self.$emit('deleteTask', self.Task_id)
                    }
                })
                .catch(function (error) {
                    console.log(error)
                })
        },
    },
}
</script>

<style scoped>
.btn-img {
    width: 20px;
    margin-right: 2px;
    margin-bottom: 2px;
}
</style>
