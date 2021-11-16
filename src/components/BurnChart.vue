<template>
<div v-if="appear=1" v-on:click="changeAppear">
  <b-button variant="success" v-on:click="DrawBurn">Draw Burn Chart</b-button>
</div>
</template>

<script >
import * as d3 from 'd3'
export default {
  name: 'BurnChart',
  data () {
    return {
      appear: 1
    }
  },
  methods: {
    changeAppear () {
      this.appear = 0
    },
    DrawBurn () {
      // 图表的宽度和高度
      var width = 600;
      var height = 600;
      // 预留给轴线的距离
      var padding = { top: 50, right: 50, bottom: 50, left: 50 };

      var dataset = [[1, 400], [2, 390], [3, 321], [4, 300], [5, 245], [6, 231], [7, 115], [8, 98], [9, 65], [10, 21]];
      var dataset1 = [[1,400],[10,21]];
      var min = d3.min(dataset, function(d) {
        return d[1];
      })
      var max = d3.max(dataset, function(d) {
        return d[1];
      })

      var xScale = d3.scaleLinear()
        .domain([1, 10])
        .range([0, width - padding.left - padding.right]);

      var yScale = d3.scaleLinear()
        .domain([0, max])
        .range([height - padding.top - padding.bottom, 0]);

      var svg = d3.select('body')
        .append('svg')
        .attr('width', width + 'px')
        .attr('height', height + 'px');

      var xAxis = d3.axisBottom()
        .scale(xScale);

      var yAxis = d3.axisLeft()
        .scale(yScale);

      svg.append('g')
        .attr('class', 'axis')
        .attr('transform', 'translate(' + padding.left + ',' + (height - padding.bottom) + ')')
        .call(xAxis);

      svg.append('g')
        .attr('class', 'axis')
        .attr('transform', 'translate(' + padding.left + ',' + padding.top + ')')
        .call(yAxis);

      var linePath = d3.line()
        .x(function(d){ return xScale(d[0]) })
        .y(function(d){ return yScale(d[1]) });

      svg.append('g')
        .append('path')
        .attr('class', 'line-path')
        .attr('transform', 'translate(' + padding.left + ',' + padding.top + ')')
        .attr('d', linePath(dataset))
        .attr('fill', 'none')
        .attr('stroke-width', 3)
        .attr('stroke', 'green');

      svg.append('g')
        .selectAll('circle')
        .data(dataset)
        .enter()
        .append('circle')
        .attr('r', 5)
        .attr('transform', function(d){
          return 'translate(' + (xScale(d[0]) + padding.left) + ',' + (yScale(d[1]) + padding.top) + ')'
        })
        .attr('fill', 'green');
      //对比线
      svg.append('g')
        .append('path')
        .attr('class', 'line-path')
        .attr('transform', 'translate(' + padding.left + ',' + padding.top + ')')
        .attr('d', linePath(dataset1))
        .attr('fill', 'none')
        .attr('stroke-width', 3)
        .attr('stroke', 'orange');

      svg.append('g')
        .selectAll('circle')
        .data(dataset1)
        .enter()
        .append('circle')
        .attr('r', 5)
        .attr('transform', function(d){
          return 'translate(' + (xScale(d[0]) + padding.left) + ',' + (yScale(d[1]) + padding.top) + ')'
        })
        .attr('fill', 'orange');
    }
  }
}

</script>

<style scoped>

</style>
