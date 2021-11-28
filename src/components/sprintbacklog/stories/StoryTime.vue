<template>
    <b-card img-top tag="article" class="mb-2">
        <b-row slot="header" style="padding-left: 15px;">
            <b-col>
                <p>Finish Time</p>
            </b-col>
        </b-row>
        <p class="card-text">
            {{ storyTime }}
        </p>
    </b-card>
</template>

<script>


export default {
    name: 'UserStoryCard',
    props: {
        story: Object,
    },
    data() {
        return {
            storyTime: String(this.story.totaltime),
        }
    },
    methods: {
        addTask(newTask) {
            this.$emit('addTask', newTask)
        },
        removeStoryAndTasks(storyId) {
            this.$emit('removeStoryAndTasks', storyId)
        },
        isDeveloper() {
            if (localStorage.getItem('teamRole') == 'Developer') {
                return true
            } else {
                return false
            }
        },
    },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Quicksand');

.text-font {
    font-family: Quicksand;
}
</style>
