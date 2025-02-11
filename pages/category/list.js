Page({
  data: {
    categories: []
  },

  onLoad: function() {
    this.loadCategories();
  },

  onShow: function() {
    this.loadCategories();
  },

  loadCategories: function() {
    const that = this;
    wx.request({
      url: 'http://localhost:8080/api/categories',
      method: 'GET',
      header: {
        'Authorization': wx.getStorageSync('token')
      },
      success: function(res) {
        that.setData({
          categories: res.data
        });
      },
      fail: function(err) {
        wx.showToast({
          title: '加载失败',
          icon: 'error'
        });
      }
    });
  },

  navigateToAdd: function() {
    wx.navigateTo({
      url: '/pages/category/add'
    });
  },

  navigateToDetail: function(e) {
    const id = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: `/pages/category/detail?id=${id}`
    });
  }
})