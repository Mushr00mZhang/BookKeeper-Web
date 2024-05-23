<template>
  <section class="bookkeeper-dashboard">
    <ElForm>
      <ElFormItem label="日期">
        <ElDatePicker
          type="daterange"
          :model-value="[search.sDate, search.eDate]"
          @update:model-value="
            ([sDate, eDate]) => {
              search.sDate = sDate;
              search.eDate = eDate;
            }
          "
          value-format="YYYY-MM-DDT00:00:00+08:00"
        />
      </ElFormItem>
    </ElForm>
    <div id="bookkeeper-dashboard-chart-daily" class="bookkeeper-dashboard-chart"></div>
  </section>
</template>
<script setup lang="ts">
import { reactive, onMounted, onActivated, onDeactivated, watch } from 'vue';
import { ElDatePicker, ElForm, ElFormItem } from 'element-plus';
import * as echarts from 'echarts';
import { Outlay } from '../Outlay/model';
const search = reactive({
  sDate: new Date(new Date().getTime() - 8 * 24 * 3600 * 1000).format('YYYY-MM-DDT00:00:00+08:00'),
  eDate: new Date(new Date().getTime() + 1 * 24 * 3600 * 1000).format('YYYY-MM-DDT00:00:00+08:00'),
});
const getXAxisData = () => {
  const dates = [];
  const sDate = new Date(search.sDate).getTime();
  const eDate = new Date(search.eDate).getTime();
  for (let i = sDate; i <= eDate; i += 24 * 3600 * 1000) {
    dates.push(new Date(i).format('YYYY-MM-DD'));
  }
  return dates;
};
let dailyChart: echarts.ECharts | null = null;
const resizeObserver = new ResizeObserver((_entries) => dailyChart?.resize());
const initDayChart = () => {
  const option: echarts.EChartsOption = {
    grid: {
      left: 64,
      right: 32,
      top: 32,
      bottom: 16,
    },
    xAxis: {
      type: 'category',
      data: getXAxisData(),
      axisLabel: {
        // showMinLabel: true,
        // showMaxLabel: true,
      },
      axisTick: {
        alignWithLabel: true,
      },
    },
    yAxis: {
      type: 'value',
    },
    tooltip: {
      show: true,
      trigger: 'axis',
    },
    series: [
      {
        type: 'line',
        data: [],
      },
    ],
  };
  const el = document.getElementById('bookkeeper-dashboard-chart-daily');
  if (el) {
    const chart = echarts.getInstanceByDom(el) || echarts.init(el, { locale: 'zh-CN' });
    dailyChart = chart;
    dailyChart.setOption(option);
    resizeObserver.observe(el);
  }
};
const getDayChartData = async () => {
  const items = await Outlay.list({ sTime: search.sDate, eTime: search.eDate });
  const dates = items.reduce<Map<string, number>>((dates, i) => {
    const date = new Date(i.time).format('YYYY-MM-DD');
    let sum = dates.get(date) || 0;
    sum += i.money;
    dates.set(date, sum);
    return dates;
  }, new Map<string, number>());
  const option: echarts.EChartsOption = {
    xAxis: {
      data: getXAxisData(),
    },
    series: [{ data: [...dates] }],
  };
  if (dailyChart) {
    dailyChart.setOption(option);
  }
};
/**
 * 轮询
 * @param func 轮询调用函数
 * @param interval 轮询间隔时间（分钟）
 */
const createLoop = async (func: Function, interval: number, init?: Function) => {
  let timeout = 0;
  const loop = async () => {
    await loop.run();
    const now = new Date().getTime();
    const next = interval * 60000 - (now % (interval * 60000));
    console.log(`loop next time: ${new Date(now + next).format('HH:mm:ss')}`);
    timeout = setTimeout(loop, next);
  };
  loop.run = async () => {
    console.log('loop running');
    await func();
  };
  loop.stop = () => {
    console.log('loop stop');
    clearTimeout(timeout);
  };
  if (init) {
    const inited = new Promise<boolean>((resolve) => {
      onMounted(async () => {
        await init();
        console.log('loop inited');
        resolve(true);
      });
      onActivated(async () => {
        await inited;
        loop();
      });
    });
  } else {
    onActivated(loop);
  }
  onDeactivated(loop.stop);
  watch(search, loop.run);
  return loop;
};
const init = async () => {
  initDayChart();
  // dayChartDataLoop.run();
};
createLoop(getDayChartData, 0.5, init);
</script>
<style lang="scss" scoped>
.bookkeeper-dashboard {
  width: 100%;
  // padding: 20px;
  overflow: hidden auto;
  .bookkeeper-dashboard-chart {
    width: 100%;
    height: 320px;
  }
}
</style>
