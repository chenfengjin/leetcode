



def quick_sort(nums):
    return quick_sort([i for i in nums if i < nums[0]]) + [i for i in nums if i == nums[0]] + quick_sort( [i for i in nums if i > nums[0]]) if len(nums)> 1 else nums


if __name__=="__main__":
    print(quick_sort([1,2,3,7,86,7,8]))