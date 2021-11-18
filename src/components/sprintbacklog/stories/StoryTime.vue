<template>
    <b-card :title="storyTime" class="mb-2 text-font">
        <b-row slot="header" style="padding-left: 15px;">
            <b-col>
                <p> Finish Time</p>
            </b-col>
        </b-row>
    </b-card>
</template>

<script>

export default {
    name: 'UserStoryCard',
    props: {
        story: Object,
    },
  data () {
        return {
            storyTime: String(this.story.description),
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
