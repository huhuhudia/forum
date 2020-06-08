module.exports = {
  devServer:{
    proxy:{
      "/api/":{
        target:'http://localhost:8888/',
        changeOrigin: true
      },
      "/apiv2/":{
        target:'http://localhost:9999/',
        changeOrigin: true
      }
    }
    
  }
}