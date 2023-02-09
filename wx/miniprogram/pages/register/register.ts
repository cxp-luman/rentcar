// pages/register/register.ts
Page({
    data: {
        licImgURL: '/resources/sedan.png' as string | undefined,
        genderIndex: 0,
        genders: ['未知', '男', '女'],
        birthDate: '1990-01-01',
    },
    onUploadLic() {
        wx.chooseMedia({
            success: res => {
                this.setData({
                    licImgURL: res.tempFiles[0].tempFilePath,
                })
            }
        })
    },
    onGenderChange(e: any) {
        this.setData({
            genderIndex: parseInt(e.detail.value),
        })
    },

})