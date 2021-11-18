<template>
    <b-container fluid>
        <b-container id="home" class="Home" fluid>
            <b-row><p></p></b-row>
            <div class="word_box">
                <p
                    style="float: left; font-size: 35px; font-weight: bolder;"
                    height="100%"
                    width="100%"
                >
                    记得改文案！！！！！
                </p>
                <p
                    style="float: left; font-size: 30px; font-weight: bold;"
                    height="100%"
                    width="100%"
                >
                    改文案！！！！！！
                </p>
            </div>
            <div class="img_box">
                <img :src="imgurl" height="100%" width="100%" />
            </div>
<!--            <b-col align-self="center">-->
<!--                <iframe-->
<!--                    width="560"-->
<!--                    height="315"-->
<!--                    src="https://www.youtube.com/embed/9leSPphzYeI"-->
<!--                    frameborder="0"-->
<!--                    allowfullscreen-->
<!--                >-->
<!--                </iframe>-->
<!--            </b-col>-->
            <br /><br /><br />
        </b-container>
    </b-container>
</template>

<script>
import teamJson from '@/assets/team.json'
import featuresJson from '@/assets/features.json'

export default {
    name: 'IndexDefault',
    data() {
        return {
            teamData: teamJson.team,
            featuresData: featuresJson.features,
            repoLink: 'https://github.com/dimitramav/DEVils-Scrum-Tool',
            imgurl: require('../../assets/login.png'),
        }
    },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=VT323');
@import url('https://fonts.googleapis.com/css?family=Quicksand');

.Home {
    background-color: #fafaff;
    width: 100%;
}

.wordcloud {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 51%;
}

.About {
    background-color: #fbfbff;
    font-family: Quicksand;
}

.Features {
    font-family: Quicksand;
    text-align: center;
    font-weight: normal;
    color: #2d2d2d;
    font-size: 13px;
}

.bottom {
    position: absolute;
    background-color: #2d2d2d;
    left: 0px;
}

.description {
    font-size: 14px;
    width: 500px;
    text-align: justify;
    word-wrap: break-word;
    font-family: Quicksand;
    padding: 8px;
    text-indent: 50px;
    text-align: justify;
    letter-spacing: 3px;
}
.word_box {
    width: 350px;
    height: 350px;
    background-color: #fff;
    border-radius: 3px;
    position: absolute;
    left: 180px;
    top: 300px;
}
.img_box {
    width: 700px;
    height: 600px;
    background-color: #fff;
    right: 120px;
    top: 80px;
    position: fixed;
}
</style>
