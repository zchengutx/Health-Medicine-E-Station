<template>
  <view class="container">
    <router-view></router-view>
  </view>
</template>

<script>
export default {
  name: 'App',
  onLaunch() {
    console.log('App Launch')
  },
  onShow() {
    console.log('App Show')
  },
  onHide() {
    console.log('App Hide')
  }
}
</script>

<style>
.container {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f5f5;
}
</style>
