/*
 * Author: robin-luo
 * Created time: 2024-03-12 15:57:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-12 16:25:50
 */
 
SELECT departmentId, MAX(salary) FROM Employee GROUP BY departmentId


SELECT Department.name Department, Employee.name Employee, salary Salary 
    FROM Employee JOIN Department ON Employee.id = Department.id
        WHERE (Employee.departmentId, salary) IN (SELECT departmentId, MAX(salary) FROM Employee GROUP BY departmentId);