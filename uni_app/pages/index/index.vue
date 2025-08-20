<template>
  <view class="index-page">
    <!-- 顶部状态栏 -->
    <view class="status-bar"></view>
    
    <!-- 顶部导航栏 -->
    <view class="header">
      <!-- 位置选择 -->
      <view class="location" @click="selectLocation">
        <text class="location-icon">📍</text>
        <text class="location-text">{{ currentLocation }}</text>
        <text class="location-arrow">▼</text>
      </view>
      
      <!-- 搜索栏 -->
      <view class="search-bar" @click="goToSearch">
        <text class="search-icon">🔍</text>
        <text class="search-placeholder">搜索药品、症状、品牌</text>
      </view>
      
      <!-- 右侧图标 -->
      <view class="header-right">
        <view class="notification" @click="goToNotification">
          <text class="notification-icon">📧</text>
        </view>
      </view>
    </view>

    <!-- 快速搜索标签 -->
    <scroll-view class="quick-search" scroll-x="true">
      <view class="search-tags">
        <view class="tag" v-for="tag in searchTags" :key="tag" @click="searchByTag(tag)">
          {{ tag }}
        </view>
      </view>
    </scroll-view>

    <!-- 主轮播图 -->
    <swiper class="main-banner" :indicator-dots="true" :autoplay="true" :interval="3000" :duration="500">
      <swiper-item v-for="banner in banners" :key="banner.id" @click="handleBannerClick(banner)">
        <image :src="banner.image" mode="aspectFill" class="banner-image"></image>
        <view class="banner-content">
          <text class="banner-title">{{ banner.title }}</text>
          <text class="banner-subtitle">{{ banner.subtitle }}</text>
          <view class="banner-tag">{{ banner.tag }}</view>
        </view>
      </swiper-item>
    </swiper>

    <!-- 特色分类 -->
    <view class="featured-categories">
      <view class="category-card" v-for="category in featuredCategories" :key="category.id" @click="goToCategory(category)">
        <view class="category-icon">
          <text class="icon">{{ category.icon }}</text>
        </view>
        <text class="category-name">{{ category.name }}</text>
        <view class="category-tag">{{ category.tag }}</view>
      </view>
    </view>

    <!-- 圆形分类导航 -->
    <view class="category-nav">
      <view class="category-item" v-for="item in categoryItems" :key="item.id" @click="goToCategory(item)">
        <view class="category-icon-circle">
          <text class="icon">{{ item.icon }}</text>
        </view>
        <text class="category-label">{{ item.name }}</text>
      </view>
    </view>

    <!-- 商品推荐 -->
    <view class="product-section">
      <view class="section-header">
        <text class="section-title">热门推荐</text>
        <text class="section-more" @click="goToMore">更多 ></text>
      </view>
      
      <scroll-view class="product-list" scroll-x="true">
        <view class="product-item" v-for="product in products" :key="product.id" @click="goToProduct(product)">
          <image :src="product.image" mode="aspectFill" class="product-image"></image>
          <view class="product-info">
            <text class="product-name">{{ product.name }}</text>
            <text class="product-desc">{{ product.description }}</text>
            <view class="product-price">
              <text class="price">¥{{ product.price }}</text>
              <text class="original-price">¥{{ product.originalPrice }}</text>
            </view>
          </view>
        </view>
      </scroll-view>
    </view>

    <!-- 新人专享 -->
    <view class="new-user-banner" @click="goToNewUserOffer">
      <view class="banner-content">
        <text class="banner-text">新人专享特惠 单品低至1元</text>
        <text class="banner-more">更多 ></text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  name: 'Index',
  data() {
    return {
      currentLocation: '东城区',
      searchTags: ['眼药水', '益生菌', '司美格鲁肽', '中暑', '流感', '减肥'],
      banners: [
        {
          id: 1,
          image: '/static/banner1.jpg',
          title: '立秋福利 健康换季',
          subtitle: '健康好物专场',
          tag: '部分满78减10/168减20'
        },
        {
          id: 2,
          image: '/static/banner2.jpg',
          title: '夏季防暑用品',
          subtitle: '清凉一夏',
          tag: '满100减20'
        }
      ],
      featuredCategories: [
        { id: 1, name: '家庭常备', icon: '👨‍👩‍👧‍👦', tag: '部分满78减10' },
        { id: 2, name: '儿童健康', icon: '⚽', tag: '部分满78减10' },
        { id: 3, name: '防暑用品', icon: '☀️', tag: '部分满78减10' },
        { id: 4, name: '皮肤用药', icon: '🩹', tag: '部分满78减10' }
      ],
      categoryItems: [
        { id: 1, name: '感冒发烧', icon: '🤧' },
        { id: 2, name: '咳嗽用药', icon: '🤒' },
        { id: 3, name: '清热解毒', icon: '🌿' },
        { id: 4, name: '皮肤用药', icon: '🧴' },
        { id: 5, name: '肠胃用药', icon: '🤢' },
        { id: 6, name: '妇科用药', icon: '🌺' },
        { id: 7, name: '男科用药', icon: '💊' },
        { id: 8, name: '儿童用药', icon: '👶' },
        { id: 9, name: '耳鼻喉药', icon: '👂' },
        { id: 10, name: '防暑避暑', icon: '🌡️' }
      ],
      products: [
        {
          id: 1,
          name: '感冒灵颗粒',
          description: '感冒发热专用',
          price: '15.8',
          originalPrice: '25.0',
          image: '/static/product1.jpg'
        },
        {
          id: 2,
          name: '维生素C片',
          description: '增强免疫力',
          price: '28.0',
          originalPrice: '35.0',
          image: '/static/product2.jpg'
        },
        {
          id: 3,
          name: '板蓝根颗粒',
          description: '清热解毒',
          price: '12.5',
          originalPrice: '18.0',
          image: '/static/product3.jpg'
        }
      ]
    }
  },
  onLoad() {
    this.loadData()
  },
  methods: {
    // 加载数据
    loadData() {
      // 这里可以调用API获取数据
      console.log('首页数据加载完成')
    },

    // 选择位置
    selectLocation() {
      uni.showToast({
        title: '位置选择功能开发中',
        icon: 'none'
      })
    },

    // 跳转到搜索页
    goToSearch() {
      uni.navigateTo({
        url: '/pages/search/search'
      })
    },

    // 跳转到通知页
    goToNotification() {
      uni.navigateTo({
        url: '/pages/notification/notification'
      })
    },

    // 按标签搜索
    searchByTag(tag) {
      uni.navigateTo({
        url: `/pages/search/search?keyword=${encodeURIComponent(tag)}`
      })
    },

    // 轮播图点击
    handleBannerClick(banner) {
      console.log('点击轮播图:', banner)
      // 根据banner类型跳转到相应页面
    },

    // 跳转到分类页
    goToCategory(category) {
      uni.navigateTo({
        url: `/pages/category/category?id=${category.id}&name=${encodeURIComponent(category.name)}`
      })
    },

    // 跳转到商品详情
    goToProduct(product) {
      uni.navigateTo({
        url: `/pages/product/product?id=${product.id}`
      })
    },

    // 跳转到更多页面
    goToMore() {
      uni.navigateTo({
        url: '/pages/product-list/product-list'
      })
    },

    // 跳转到新人专享
    goToNewUserOffer() {
      uni.navigateTo({
        url: '/pages/new-user/new-user'
      })
    }
  }
}
</script>

<style scoped>
.index-page {
  min-height: 100vh;
  background-color: #f5f5f5;
}

.status-bar {
  height: 44rpx;
  background-color: #e74c3c;
}

.header {
  display: flex;
  align-items: center;
  padding: 20rpx 30rpx;
  background: linear-gradient(135deg, #e74c3c 0%, #c0392b 100%);
  color: white;
}

.location {
  display: flex;
  align-items: center;
  margin-right: 20rpx;
  font-size: 28rpx;
}

.location-icon {
  font-size: 24rpx;
  margin-right: 8rpx;
}

.location-arrow {
  font-size: 20rpx;
  margin-left: 8rpx;
}

.search-bar {
  flex: 1;
  display: flex;
  align-items: center;
  background: white;
  border-radius: 50rpx;
  padding: 15rpx 25rpx;
  margin: 0 20rpx;
}

.search-icon {
  font-size: 28rpx;
  color: #999;
  margin-right: 15rpx;
}

.search-placeholder {
  font-size: 28rpx;
  color: #999;
}

.header-right {
  display: flex;
  align-items: center;
}

.notification-icon {
  font-size: 36rpx;
}

.quick-search {
  background: linear-gradient(135deg, #e74c3c 0%, #c0392b 100%);
  padding: 20rpx 0;
  white-space: nowrap;
}

.search-tags {
  display: flex;
  padding: 0 30rpx;
}

.tag {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  padding: 10rpx 20rpx;
  border-radius: 25rpx;
  font-size: 24rpx;
  margin-right: 20rpx;
  white-space: nowrap;
}

.main-banner {
  height: 300rpx;
  margin: 20rpx 30rpx;
  border-radius: 20rpx;
  overflow: hidden;
}

.banner-image {
  width: 100%;
  height: 100%;
}

.banner-content {
  position: absolute;
  top: 30rpx;
  left: 30rpx;
  color: white;
}

.banner-title {
  display: block;
  font-size: 24rpx;
  margin-bottom: 10rpx;
  opacity: 0.9;
}

.banner-subtitle {
  display: block;
  font-size: 36rpx;
  font-weight: bold;
  margin-bottom: 15rpx;
}

.banner-tag {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  padding: 8rpx 16rpx;
  border-radius: 15rpx;
  font-size: 22rpx;
}

.featured-categories {
  display: flex;
  padding: 0 30rpx;
  margin-bottom: 30rpx;
}

.category-card {
  flex: 1;
  background: white;
  margin: 0 10rpx;
  padding: 30rpx 20rpx;
  border-radius: 15rpx;
  text-align: center;
  position: relative;
}

.category-icon {
  margin-bottom: 15rpx;
}

.icon {
  font-size: 48rpx;
}

.category-name {
  display: block;
  font-size: 26rpx;
  color: #333;
  margin-bottom: 10rpx;
}

.category-tag {
  background: #e74c3c;
  color: white;
  padding: 6rpx 12rpx;
  border-radius: 10rpx;
  font-size: 20rpx;
  position: absolute;
  top: 15rpx;
  right: 15rpx;
}

.category-nav {
  background: white;
  padding: 30rpx;
  margin-bottom: 20rpx;
}

.category-item {
  display: inline-block;
  width: 20%;
  text-align: center;
  margin-bottom: 30rpx;
}

.category-icon-circle {
  width: 80rpx;
  height: 80rpx;
  background: #f8f9fa;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 15rpx;
}

.category-icon-circle .icon {
  font-size: 36rpx;
}

.category-label {
  font-size: 24rpx;
  color: #666;
}

.product-section {
  background: white;
  padding: 30rpx;
  margin-bottom: 20rpx;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
}

.section-title {
  font-size: 32rpx;
  font-weight: bold;
  color: #333;
}

.section-more {
  font-size: 26rpx;
  color: #999;
}

.product-list {
  white-space: nowrap;
}

.product-item {
  display: inline-block;
  width: 300rpx;
  margin-right: 20rpx;
  background: #f8f9fa;
  border-radius: 15rpx;
  overflow: hidden;
}

.product-image {
  width: 100%;
  height: 200rpx;
}

.product-info {
  padding: 20rpx;
}

.product-name {
  display: block;
  font-size: 28rpx;
  color: #333;
  margin-bottom: 10rpx;
  white-space: normal;
  line-height: 1.4;
}

.product-desc {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-bottom: 15rpx;
  white-space: normal;
}

.product-price {
  display: flex;
  align-items: center;
}

.price {
  font-size: 32rpx;
  color: #e74c3c;
  font-weight: bold;
  margin-right: 10rpx;
}

.original-price {
  font-size: 24rpx;
  color: #999;
  text-decoration: line-through;
}

.new-user-banner {
  background: #e74c3c;
  margin: 0 30rpx 30rpx;
  border-radius: 15rpx;
  padding: 30rpx;
}

.banner-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white;
}

.banner-text {
  font-size: 28rpx;
  font-weight: bold;
}

.banner-more {
  font-size: 26rpx;
  opacity: 0.8;
}
</style>
