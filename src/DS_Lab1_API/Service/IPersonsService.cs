using DS_Lab1.CommunicationElements;
using DS_Lab1.Models;

namespace DS_Lab1.Service
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
