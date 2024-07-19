/// [Leetcode](https://leetcode.cn) solution
pub struct Solution;

use std::collections::HashMap;

impl Solution {
    /// [1. Tow Sum](https://leetcode.cn/problems/two-sum/description)
    ///
    /// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
    ///
    /// You may assume that each input would have exactly one solution, and you may not use the same element twice.
    ///
    /// You can return the answer in any order.
    ///
    ///  
    ///
    /// Example 1:
    ///
    //      Input: nums = [2,7,11,15], target = 9
    ///     Output: [0,1]
    ///     Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
    ///
    /// Example 2:
    ///
    ///     Input: nums = [3,2,4], target = 6
    ///     Output: [1,2]
    ///
    /// Example 3:
    ///
    ///     Input: nums = [3,3], target = 6
    ///     Output: [0,1]
    ///     
    ///  
    ///
    /// Constraints:
    ///
    ///     2 <= nums.length <= 104
    ///     -109 <= nums[i] <= 109
    ///     -109 <= target <= 109
    ///     Only one valid answer exists.
    ///
    ///  
    /// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut indices = HashMap::new();
        for (i, v) in nums.iter().enumerate() {
            if let Some(&idx) = indices.get(&(target - v)) {
                return vec![i as i32, idx as i32];
            };
            indices.insert(v, i);
        }
        return vec![];
    }

    /// [49. Group Anagrams](https://leetcode.cn/problems/group-anagrams/description/)
    ///
    /// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
    ///
    /// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: strs = ["eat","tea","tan","ate","nat","bat"]
    /// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
    ///
    /// Example 2:
    ///
    /// Input: strs = [""]
    /// Output: [[""]]
    ///
    /// Example 3:
    ///
    /// Input: strs = ["a"]
    /// Output: [["a"]]
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     1 <= strs.length <= 104
    ///     0 <= strs[i].length <= 100
    ///     strs[i] consists of lowercase English letters.
    ///
    ///
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut matchers = HashMap::new();
        for str in strs.iter() {
            let mut matcher = [0; 26];
            for e in str.chars() {
                matcher[e as usize - 'a' as usize] += 1;
            }
            let group = matchers.entry(matcher).or_insert(vec![]);
            group.push(str.to_string());
        }

        let mut result: Vec<Vec<String>> = Vec::with_capacity(matchers.len());
        for (_, group) in matchers {
            result.push(group);
        }
        return result;
    }

    /// [128. Longest Consecutive Sequence](https://leetcode.cn/problems/longest-consecutive-sequence/description/)
    ///
    /// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
    ///
    /// You must write an algorithm that runs in O(n) time.
    ///
    /// Example 1:
    ///
    /// Input: nums = [100,4,200,1,3,2]
    /// Output: 4
    /// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
    ///
    /// Example 2:
    ///
    /// Input: nums = [0,3,7,2,5,8,4,6,0,1]
    /// Output: 9
    ///
    /// Constraints:
    ///
    ///	0 <= nums.length <= 10^5
    ///	-10^9 <= nums[i] <= 10^9
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut exist: HashMap<i32, bool> = HashMap::new();
        for num in nums {
            exist.insert(num, true);
        }

        let mut longest = 0;
        for (&num, _) in exist.iter() {
            if let Some(&ok) = exist.get(&(num - 1)) {
                if ok {
                    continue;
                }
            }

            let mut length = 1;
            let mut n = num;
            while let Some(&ok) = exist.get(&(n + 1)) {
                if !ok {
                    break;
                }
                length += 1;
                n += 1;
            }
            longest = std::cmp::max(longest, length)
        }
        return longest;
    }

    /// [283. Move Zeroes](https://leetcode.cn/problems/move-zeroes/description/)
    ///
    /// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
    ///
    /// Note that you must do this in-place without making a copy of the array.
    ///
    /// Example 1:
    ///
    /// Input: nums = [0,1,0,3,12]
    /// Output: [1,3,12,0,0]
    ///
    /// Example 2:
    ///
    /// Input: nums = [0]
    /// Output: [0]
    ///
    /// Constraints:
    ///
    ///	1 <= nums.length <= 10^4
    ///	-2^31 <= nums[i] <= 2^31 - 1
    ///
    /// Follow up: Could you minimize the total number of operations done?
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut tail = 0;
        for head in 0..nums.len() {
            if nums[head] != 0 {
                nums.swap(tail, head);
                tail += 1;
            }
        }
    }

    /// [11. Container with most water](https://leetcode.cn/problems/container-with-most-water/description/)
    ///
    /// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
    ///
    /// Find two lines that together with the x-axis form a container, such that the container contains the most water.
    ///
    /// Return the maximum amount of water a container can store.
    ///
    /// Notice that you may not slant the container.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// ```txt
    ///  8|                    |                |              
    ///  7|                    |                |              |
    ///  6|          |         |          |     |              |
    ///  5|          |         |          |     |       |      |
    ///  4|          |         |          |     |       |      |
    ///  3|          |         |          |     |       |      |
    ///  2|          |         |    |     |     |       |      |
    ///  1|    |     |         |    |     |     |       |      |
    ///   -------------------------------------------------------
    ///   0    1     2         3    4     5     6       7      8
    /// ```
    ///
    /// Input: height = [1,8,6,2,5,4,8,3,7]
    ///
    /// Output: 49
    ///
    /// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
    ///
    /// Example 2:
    ///
    /// Input: height = [1,1]
    /// Output: 1
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     n == height.length
    ///     2 <= n <= 105
    ///     0 <= height[i] <= 104
    ///
    ///
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut i = 0;
        let mut j = height.len() - 1;
        let mut max_area = 0;

        while j > i {
            let area = height[i].min(height[j]) * (j - i) as i32;
            max_area = max_area.max(area);
            if height[i] > height[j] {
                j -= 1;
            } else {
                i += 1;
            }
        }
        return max_area;
    }

    /// [15. 3Sum](https://leetcode.cn/problems/3sum/description/)
    ///
    /// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
    ///
    /// Notice that the solution set must not contain duplicate triplets.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: nums = [-1,0,1,2,-1,-4]
    /// Output: [[-1,-1,2],[-1,0,1]]
    /// Explanation:
    /// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
    /// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
    /// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
    /// The distinct triplets are [-1,0,1] and [-1,-1,2].
    /// Notice that the order of the output and the order of the triplets does not matter.
    ///
    /// Example 2:
    ///
    /// Input: nums = [0,1,1]
    /// Output: []
    /// Explanation: The only possible triplet does not sum up to 0.
    ///
    /// Example 3:
    ///
    /// Input: nums = [0,0,0]
    /// Output: [[0,0,0]]
    /// Explanation: The only possible triplet sums up to 0.
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     3 <= nums.length <= 3000
    ///     -105 <= nums[i] <= 105
    ///     
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort();
        let mut result = Vec::new();
        for k in 0..(nums.len() - 2) {
            if nums[k] > 0 {
                break;
            }
            if k > 0 && nums[k] == nums[k - 1] {
                continue;
            }

            let mut i = k + 1;
            let mut j = nums.len() - 1;
            while i < j {
                let sum = nums[k] + nums[i] + nums[j];
                if sum == 0 {
                    result.push(vec![nums[k], nums[i], nums[j]])
                }
                if sum > 0 {
                    j = j - 1;
                    while i < j && nums[j] == nums[j + 1] {
                        j -= 1;
                    }
                } else {
                    i = i + 1;
                    while i < j && nums[i] == nums[i - 1] {
                        i += 1;
                    }
                }
            }
        }
        return result;
    }

    /// [42. Trapping Rain Water](https://leetcode.cn/problems/trapping-rain-water/description/)
    ///
    /// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
    ///
    ///  
    ///
    /// Example 1:
    ///
    /// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
    /// Output: 6
    /// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
    ///
    /// Example 2:
    ///
    /// Input: height = [4,2,0,3,2,5]
    /// Output: 9
    ///
    ///  
    ///
    /// Constraints:
    ///
    ///     n == height.length
    ///     1 <= n <= 2 * 104
    ///     0 <= height[i] <= 105
    ///
    ///
    pub fn trap(height: Vec<i32>) -> i32 {
        if height.len() == 0 {
            return 0;
        }

        let mut total = 0;
        let mut left = 0;
        let mut right = height.len() - 1;
        let mut left_max = 0;
        let mut right_max = 0;

        while left < right {
            left_max = left_max.max(height[left]);
            right_max = right_max.max(height[right]);

            if height[left] < height[right] {
                total += left_max - height[left];
                left += 1;
            } else {
                total += right_max - height[right];
                right -= 1;
            }
        }
        return total;
    }

    /// [3. Longest Substring Without Repeating Characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/)
    ///
    /// Given a string s, find the length of the longest
    /// substring
    /// without repeating characters.
    ///
    ///
    ///
    /// Example 1:
    ///
    /// Input: s = "abcabcbb"
    /// Output: 3
    /// Explanation: The answer is "abc", with the length of 3.
    ///
    /// Example 2:
    ///
    /// Input: s = "bbbbb"
    /// Output: 1
    /// Explanation: The answer is "b", with the length of 1.
    ///
    /// Example 3:
    ///
    /// Input: s = "pwwkew"
    /// Output: 3
    /// Explanation: The answer is "wke", with the length of 3.
    /// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
    ///
    ///
    ///
    /// Constraints:
    ///
    ///     0 <= s.length <= 5 * 104
    ///     s consists of English letters, digits, symbols and spaces.
    ///
    pub fn length_of_longest_substring(s: String) -> i32 {
        if s.len() <= 1 {
            return s.len() as i32;
        }

        let mut left = 0;
        let mut right = 0;
        let mut longest = 0;
        let mut m: HashMap<char, usize> = HashMap::new();
        let chars: Vec<char> = s.chars().collect();
        while left <= right && right < chars.len() {
            if let Some(&idx) = &m.get(&chars[right]) {
                for i in left..(idx) {
                    m.remove(&chars[i]);
                }
                left = idx + 1;
            } else {
                longest = longest.max(right - left + 1);
            }
            m.entry(chars[right]).or_insert(right);
            right += 1;
        }
        return longest as i32;
    }
}
