Page({
  data: {
    userInfo: null
  },

  onLoad: function() {
    this.loadUserInfo();
  },

  onShow: function() {
    this.loadUserInfo();
  },

  loadUserInfo: function() {
    const that = this;
    const token = wx.getStorageSync('token');
    
    if (!token) {
      return;
    }

    wx.request({
      url: 'http://localhost:8080/api/user/profile',
      method: 'GET',
      header: {
        'Authorization': token
      },
      success: function(res) {
        that.setData({
          userInfo: res.data
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

  login: function() {
    const that = this;
    wx.login({
      success: function(res) {
        if (res.code) {
          wx.request({
            url: 'http://localhost:8080/api/auth/login',
            method: 'POST',
            data: {
              code: res.code
            },
            success: function(res) {
              wx.setStorageSync('token', res.data.token);
              that.loadUserInfo();
            },
            fail: function(err) {
              wx.showToast({
                title: '登录失败',
                icon: 'error'
              });
            }
          });
        }
      }
    });
  },

  navigateToAssets: function() {
    wx.switchTab({
      url: '/pages/asset/list'
    });
  },

  navigateToCategories: function() {
    wx.switchTab({
      url: '/pages/category/list'
    });
  },

  navigateToSettings: function() {
    wx.navigateTo({
      url: '/pages/user/settings'
    });
  }
})