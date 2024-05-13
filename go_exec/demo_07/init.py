import matplotlib.pyplot as plt
import numpy as np

# 定义约束条件
x1 = np.linspace(0, 8, 100)
x2_1 = (24 - 3*x1)/4  # x2的约束条件 3x1+4x2 <= 24
x2_2 = 8 - x1  # x2的约束条件 x1+x2+x3 <= 8
x2_3 = np.zeros_like(x1)

# 绘制图形
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(12, 4))
ax1.plot(x1, x2_1, label='3x1+4x2=24')
ax1.plot(x1, x2_2, label='x1+x2+x3=8')
ax1.plot(x1, x2_3, label='x2>=0')
ax1.fill_betweenx(x2_1, 0, x1, alpha=0.1)
ax1.fill_betweenx(x2_2, 0, x1, where=x2_2>=0, alpha=0.1)
ax1.set_xlim(0, 8)
ax1.set_ylim(0, 6)
ax1.set_xlabel('$x_1$')
ax1.set_ylabel('$x_2$')
ax1.legend()

x2 = np.linspace(0, 8, 100)
x3_1 = (24 - 3*x1 - 4*x2)/2  # x3的约束条件 3x1+4x2+2x3 <= 24
x3_2 = 8 - x1 - x2  # x3的约束条件 x1+x2+x3 <= 8
x3_3 = np.zeros_like(x2)

# 绘制图形
ax2.plot(x2, x3_1, label='3x1+4x2+2x3=24')
ax2.plot(x2, x3_2, label='x1+x2+x3=8')
ax2.plot(x2, x3_3, label='x3>=0')
ax2.fill_betweenx(x3_1, 0, x2, alpha=0.1)
ax2.fill_betweenx(x3_2, 0, x2, where=x3_2>=0, alpha=0.1)
ax2.set_xlim(0, 8)
ax2.set_ylim(0, 6)
ax2.set_xlabel('$x_2$')
ax2.set_ylabel('$x_3$')
ax2.legend()

plt.show()
