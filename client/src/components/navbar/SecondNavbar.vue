<template>
    <b-navbar id="second-nav" fixed="top" class="subnav" type="light">
        <b-navbar-nav v-if="logout_prop === true" class="mx-auto">
            <b-nav-item href="#home">
                <p class="anchors"></p>
            </b-nav-item>
            <b-col></b-col>
            <b-nav-item href="#features">
                <p class="anchors"></p>
            </b-nav-item>
            <b-col></b-col>
            <b-nav-item href="#about">
                <p class="anchors"></p>
            </b-nav-item>
        </b-navbar-nav>
        <b-navbar-nav v-if="logout_prop === false" class="mx-auto" >
            <b-nav-item
                :to="{
                    name: 'OldSprints',
                    params: { id: $route.params.id },
                }"
            >
                <p class="anchors">Sprint backlog</p>
            </b-nav-item>
            <b-col></b-col>
            <b-nav-item
                :to="{
                    name: 'ProjectPageOverview',
                    params: { id: $route.params.id },
                }"
            >
                <p class="anchors">Overview</p>
            </b-nav-item>
            <b-col></b-col>
            <b-nav-item
                :to="{
                    name: 'Backlog',
                    params: { id: $route.params.id },
                }"
            >
                <p class="anchors">Product Backlog</p>
            </b-nav-item>
            <b-col></b-col>
            <b-nav-item-dropdown  class="anchors" text="Meeting" toggle-class="text-light">
                <b-nav-item
                    :to="{
                        name: 'Meeting',
                        params: { id: $route.params.id },
                    }"
                >
                    <p>Booking</p>
                </b-nav-item>
              <b-nav-item
                :to="{
                        name: 'Calendar',
                        params: { id: $route.params.id },
                    }"
              >
                <p>Details</p>
              </b-nav-item>
            </b-nav-item-dropdown>
        </b-navbar-nav>
    </b-navbar>
</template>

<script>
export default {
    name: 'SecondNavbar',
    props: {
        logout_prop: Boolean,
    },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Merienda');
@import url('https://fonts.googleapis.com/css?family=Quicksand');

.text-enhancement2 {
    font-family: 'Quicksand', cursive;
    font-size: 18px;
    font-weight: bold;
    color: white;
    margin-top: 2px;
    margin-bottom: 5px;
}

.anchors {
    font-family: 'Merienda', cursive;
    color: white;
    font-style: italic;
    font-weight: bold;
    margin-top: 5px;
    margin-bottom: 5px;
}

.subnav {
    background-color: #494965 !important;
    top: 70px;
    padding-top: 4px;
    padding-bottom: 2px;
    line-height: 2px;
}
.menu {
    width: 30px;
    position: relative;
    bottom: 0.5px;
}
</style>
