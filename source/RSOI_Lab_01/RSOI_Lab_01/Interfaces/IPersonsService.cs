using RSOI_Lab_01.Models;

namespace RSOI_Lab_01.Interfaces
{
    public interface IPersonsService
    {
        Task<IEnumerable<PersonResponse>> GetAllAsync();
        Task<PersonResponse?> GetAsync(int id);
        Task<int?> CreateAsync(PersonRequest person);
        Task<PersonResponse?> UpdateAsync(int id, PersonRequest person);
        Task DeleteAsync(int id);
    }
}

