Page({
  data: {
    assets: []
  },

  onLoad: function() {
    this.loadAssets();
  },

  onShow: function() {
    this.loadAssets();
  },

  loadAssets: function() {
    const that = this;
    wx.request({
      url: 'http://localhost:8080/api/assets',
      method: 'GET',
      header: {
        'Authorization': wx.getStorageSync('token')
      },
      success: function(res) {
        that.setData({
          assets: res.data
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
      url: '/pages/asset/add'
    });
  },

  navigateToDetail: function(e) {
    const id = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: `/pages/asset/detail?id=${id}`
    });
  }
})