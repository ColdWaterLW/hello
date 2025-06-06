document.addEventListener('DOMContentLoaded', () => {
    const ctaButton = document.getElementById('ctaButton');

    if (ctaButton) {
        ctaButton.addEventListener('click', () => {
            alert('您点击了“了解更多”按钮！');
            // 这里可以添加更复杂的逻辑，例如滚动到某个部分，或者显示一个模态框
            // 示例：平滑滚动到“关于我们”部分
            const aboutSection = document.getElementById('about');
            if (aboutSection) {
                aboutSection.scrollIntoView({
                    behavior: 'smooth'
                });
            }
        });
    }

    // 导航栏平滑滚动效果 (可选)
    document.querySelectorAll('nav a').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();

            const targetId = this.getAttribute('href').substring(1);
            const targetSection = document.getElementById(targetId);

            if (targetSection) {
                targetSection.scrollIntoView({
                    behavior: 'smooth'
                });
            }
        });
    });

    // 可以在这里添加更多的交互逻辑，例如：
    // - 表单验证
    // - 动态内容加载
    // - 轮播图
    // - 动画效果
});